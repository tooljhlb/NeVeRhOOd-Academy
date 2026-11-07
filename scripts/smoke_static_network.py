#!/usr/bin/env python3
"""One-shot localhost network smoke for the Academy static site. Stdlib only."""
from __future__ import annotations
import contextlib
from functools import partial
from http.server import SimpleHTTPRequestHandler, ThreadingHTTPServer
from pathlib import Path
from threading import Thread
from urllib.request import urlopen

ROOT=Path(__file__).resolve().parents[1]
SITE=ROOT/"site"
PATHS={
    "/": "NeVeRhOOd Academy",
    "/track/foundations/index.html":"AI Engineering Foundations",
    "/learn/ai-agents-101/index.html":"AI Agents 101",
    "/learn/protocols-101/index.html":"Protocols 101",
    "/learn/evaluation-101/index.html":"Model Evaluation 101",
    "/captain/index.html":"Captain",
    "/about/index.html":"NeVeRhOOd",
}
class QuietHandler(SimpleHTTPRequestHandler):
    def log_message(self,*args): pass

def main()->None:
    handler=partial(QuietHandler,directory=str(SITE))
    server=ThreadingHTTPServer(("127.0.0.1",0),handler)
    thread=Thread(target=server.serve_forever,daemon=True)
    thread.start()
    try:
        base=f"http://127.0.0.1:{server.server_address[1]}"
        for url,needle in PATHS.items():
            with contextlib.closing(urlopen(base+url,timeout=5)) as reply:
                assert reply.status==200,(url,reply.status)
                assert "text/html" in reply.headers.get("Content-Type",""),url
                document=reply.read(1_000_000).decode("utf-8")
                assert needle in document,(url,"expected title missing")
        print(f"ACADEMY_HTTP_SMOKE_PASS pages={len(PATHS)} localhost_only=true")
    finally:
        server.shutdown()
        server.server_close()
        thread.join(timeout=2)

if __name__=="__main__":
    main()
