#!/usr/bin/env python3
"""Validate Academy catalog records without claiming unverified credits or certificates.

Uses Python standard library only; suitable for CI, Manus import gates, and local builds.
"""
from __future__ import annotations
import argparse
import csv
import json
import sys
from collections import Counter
from pathlib import Path
from urllib.parse import urlsplit

ROOT = Path(__file__).resolve().parents[1]
REQUIRED = ("id", "title", "provider", "category", "url", "credential_status", "link_status")
BLOCKED_SCHEMES = {"file", "javascript", "data"}
PRIVATE_HOSTS = {"localhost", "127.0.0.1", "::1"}


def validate(path: Path) -> dict:
    errors: list[str] = []
    warnings: list[str] = []
    rows: list[dict] = []
    seen_ids: dict[str, int] = {}
    seen_links: dict[str, int] = {}
    if not path.is_file():
        return {"ok": False, "rows": 0, "distinct_urls": 0,
                "errors": [f"Catalog missing: {path}"], "warnings": []}
    with path.open("r", newline="", encoding="utf-8-sig") as fh:
        reader = csv.DictReader(fh)
        columns = reader.fieldnames or []
        for field in REQUIRED:
            if field not in columns:
                errors.append(f"Missing column: {field}")
        if errors:
            return {"ok": False, "rows": 0, "distinct_urls": 0, "errors": errors, "warnings": []}
        for line, row in enumerate(reader, start=2):
            rows.append(row)
            for field in ("id", "title", "provider", "url"):
                if not (row.get(field) or "").strip():
                    errors.append(f"Line {line}: {field} is required")
            resource_id = (row.get("id") or "").strip()
            if resource_id in seen_ids:
                errors.append(f"Line {line}: duplicate id {resource_id!r} (first at line {seen_ids[resource_id]})")
            else:
                seen_ids[resource_id] = line
            uri = (row.get("url") or "").strip()
            try:
                parsed = urlsplit(uri)
                if parsed.scheme.lower() in BLOCKED_SCHEMES or parsed.scheme.lower() != "https":
                    errors.append(f"Line {line}: URL must be HTTPS")
                if not parsed.hostname or parsed.hostname.lower() in PRIVATE_HOSTS or parsed.username or parsed.password:
                    errors.append(f"Line {line}: invalid/untrusted URL host")
                key = (parsed.hostname or "").casefold() + parsed.path.rstrip("/").casefold()
                if key in seen_links:
                    warnings.append(f"Line {line}: repeats destination from line {seen_links[key]}")
                else:
                    seen_links[key] = line
            except ValueError:
                errors.append(f"Line {line}: malformed URL")
            if (row.get("credential_status") or "").lower() not in {
                "unknown", "unverified", "not-confirmed", "check-provider",
                "none", "provider-certificate", "academy-completion"
            }:
                warnings.append(f"Line {line}: credential status needs issuer verification")
            if (row.get("link_status") or "").lower() not in {
                "verified", "unverified", "not-confirmed", "broken", "reviewing"
            }:
                warnings.append(f"Line {line}: link status unknown")
    return {
        "ok": not errors,
        "rows": len(rows),
        "distinct_urls": len(seen_links),
        "by_category": dict(sorted(Counter(r["category"].strip() for r in rows).items())),
        "errors": errors,
        "warnings": warnings,
        "note": "This validation does NOT check whether providers award certificates or academic credit.",
    }


def main() -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument("--catalog", type=Path, default=ROOT / "data" / "courses.csv")
    parser.add_argument("--output", type=Path)
    args = parser.parse_args()
    result = validate(args.catalog)
    out = json.dumps(result, indent=2, ensure_ascii=False) + "\n"
    if args.output:
        args.output.parent.mkdir(parents=True, exist_ok=True)
        args.output.write_text(out, encoding="utf-8")
    print(out)
    return 0 if result["ok"] else 1


if __name__ == "__main__":
    sys.exit(main())
