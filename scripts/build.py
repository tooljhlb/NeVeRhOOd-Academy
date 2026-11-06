#!/usr/bin/env python3
"""Build a searchable static Academy course catalogue from reviewed CSV data."""
import csv
import html
import json
from pathlib import Path
from urllib.parse import urlparse

ROOT = Path(__file__).resolve().parents[1]
SOURCE = ROOT / "data" / "courses.csv"
OUT = ROOT / "site"
ALLOWED_SCHEMES = {"https"}
def canonical(url):
    url = url.strip()
    p = urlparse(url)
    if p.scheme not in ALLOWED_SCHEMES or not p.hostname:
        raise ValueError("Invalid course URL: " + url)
    return p._replace(fragment="", query="", path=p.path.rstrip("/") or "/").geturl()
def build():
    rows = []
    seen_urls = set()
    with SOURCE.open(newline="", encoding="utf-8-sig") as fp:
        for r in csv.DictReader(fp):
            r = {k: (v or "").strip() for k,v in r.items()}
            u = canonical(r["url"])
            if u in seen_urls:
                continue
            seen_urls.add(u)
            r["url"] = u
            rows.append(r)
    OUT.mkdir(parents=True, exist_ok=True)
    (OUT / "courses.json").write_text(json.dumps(rows, indent=2, ensure_ascii=False) + "\n", encoding="utf-8")
    cards = []
    for r in rows:
        title, provider, category, url = [html.escape(r[k], quote=True) for k in ("title", "provider", "category", "url")]
        search = html.escape(" ".join((r["title"],r["provider"],r["category"])).lower(), quote=True)
        cards.append(f'<article data-search="{search}"><h2>{title}</h2><p>{provider} · {category}</p><p>Certificate: {html.escape(r["credential_status"])} · Link: {html.escape(r["link_status"])}</p><a href="{url}" target="_blank" rel="noopener noreferrer">Visit original provider</a></article>')
    template = (ROOT / "site" / "template.html").read_text(encoding="utf-8")
    (OUT / "index.html").write_text(template.replace("<!--COURSE_CARDS-->", "\n".join(cards)), encoding="utf-8")
    print(f"Built {len(rows)} unique course records")
if __name__ == "__main__":
    build()
