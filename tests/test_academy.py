"""Deterministic, dependency-free Academy source tests."""
import csv
import importlib.util
import tempfile
import unittest
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
spec = importlib.util.spec_from_file_location("validate_catalog", ROOT / "scripts" / "validate_catalog.py")
validate_catalog = importlib.util.module_from_spec(spec)
assert spec.loader
spec.loader.exec_module(validate_catalog)


class AcademyTests(unittest.TestCase):
    def test_pilot_catalog_is_structurally_valid(self):
        report = validate_catalog.validate(ROOT / "data" / "courses.csv")
        self.assertTrue(report["ok"], report["errors"])
        self.assertGreaterEqual(report["rows"], 4)

    def test_duplicate_ids_refused(self):
        with tempfile.TemporaryDirectory() as tmp:
            file = Path(tmp) / "data.csv"
            with file.open("w", encoding="utf-8", newline="") as fh:
                writer = csv.DictWriter(fh, fieldnames=validate_catalog.REQUIRED)
                writer.writeheader()
                for title in ("A", "B"):
                    writer.writerow(dict(id="same", title=title, provider="Example", category="AI",
                                         url=f"https://example.com/{title}", credential_status="unknown",
                                         link_status="unverified"))
            self.assertFalse(validate_catalog.validate(file)["ok"])

    def test_local_or_http_url_refused(self):
        with tempfile.TemporaryDirectory() as tmp:
            file = Path(tmp) / "data.csv"
            with file.open("w", encoding="utf-8", newline="") as fh:
                writer = csv.DictWriter(fh, fieldnames=validate_catalog.REQUIRED)
                writer.writeheader()
                writer.writerow(dict(id="a", title="A", provider="Example", category="AI",
                                     url="http://localhost/private", credential_status="unknown",
                                     link_status="unverified"))
            self.assertFalse(validate_catalog.validate(file)["ok"])

    def test_external_credit_claims_not_embedded_in_course_schema(self):
        schema = (ROOT / "courses" / "schema.json").read_text(encoding="utf-8")
        self.assertIn('"verified_credits"', schema)
        self.assertIn('"official_evidence_url"', schema)

    def test_original_lesson_exists_with_five_quiz_questions(self):
        src = (ROOT / "site" / "learn" / "ai-agents-101" / "index.html").read_text(encoding="utf-8")
        self.assertEqual(src.count("<fieldset>"), 5)
        self.assertIn("non-accredited", src.lower())

    def test_foundations_track_contains_all_three_lessons(self):
        page = (ROOT / "site" / "track" / "foundations" / "index.html").read_text(encoding="utf-8")
        for slug in ("ai-agents-101", "protocols-101", "evaluation-101"):
            self.assertIn(f"learn/{slug}/index.html", page)

    def test_two_new_lessons_include_five_practice_questions_each(self):
        for slug in ("protocols-101", "evaluation-101"):
            page = (ROOT / "site" / "learn" / slug / "index.html").read_text(encoding="utf-8")
            self.assertEqual(page.count("<fieldset>"), 5, slug)
            self.assertIn("not an accredited credential", page.lower(), slug)
            self.assertIn('type="submit"', page, slug)

    def test_homepage_template_features_foundations_track(self):
        page = (ROOT / "site" / "template.html").read_text(encoding="utf-8")
        self.assertIn('href="track/foundations/index.html"', page)

    def test_academy_template_links_to_own_lesson(self):
        src = (ROOT / "site" / "template.html").read_text(encoding="utf-8")
        self.assertIn("learn/ai-agents-101/index.html", src)


if __name__ == "__main__":
    unittest.main()
