// Dependency-free behavioral checks for the Academy workspace UI.
const { test } = require('node:test');
const assert = require('node:assert/strict');
const fs = require('node:fs');
const vm = require('node:vm');
const path = require('node:path');
const html = fs.readFileSync(path.join(__dirname, '../site/workspace/index.html'), 'utf8');
const script = html.match(/<script>([\s\S]*?)<\/script>/)?.[1];
if (!script) throw new Error('workspace script missing');

function harness() {
  const nodes = new Map();
  const listeners = new Map();
  function element(id) {
    if (!nodes.has(id)) {
      const attributes = new Map();
      const handlers = new Map();
      const classes = new Set();
      nodes.set(id, {
        id, dataset: {}, hidden: false,
        getAttribute(name) { return attributes.get(name) ?? null; },
        setAttribute(name, value) { attributes.set(name, String(value)); },
        addEventListener(name, callback) { handlers.set(name, callback); },
        click() { const callback = handlers.get('click'); if (callback) callback({ type: 'click' }); },
        classList: {
          contains(name) { return classes.has(name); },
          add(name) { classes.add(name); },
          remove(name) { classes.delete(name); },
          toggle(name, on) {
            if (on === undefined ? !classes.has(name) : on) { classes.add(name); return true; }
            classes.delete(name); return false;
          },
        },
      });
    }
    return nodes.get(id);
  }
  const lessonButtons = ['foundations', 'agents', 'evaluation', 'untrusted'].map(name => {
    const button = element('lesson-' + name);
    button.dataset.lesson = name;
    return button;
  });
  const document = {
    getElementById: element,
    querySelectorAll(selector) {
      assert.equal(selector, '[data-lesson]');
      return lessonButtons;
    },
    addEventListener(name, callback) { listeners.set(name, callback); },
  };
  vm.runInNewContext(script, { document, Object }, { timeout: 1000 });
  return { element, listeners };
}

test('drawer toggle synchronizes visible state and ARIA', () => {
  const { element } = harness();
  const menu = element('menu'), nav = element('nav'), overlay = element('overlay');
  menu.click();
  assert.equal(nav.classList.contains('open'), true);
  assert.equal(overlay.classList.contains('open'), true);
  assert.equal(menu.getAttribute('aria-expanded'), 'true');
  menu.click();
  assert.equal(nav.classList.contains('open'), false);
  assert.equal(menu.getAttribute('aria-expanded'), 'false');
});

test('overlay and Escape close the navigation drawer', () => {
  const { element, listeners } = harness();
  element('menu').click();
  element('overlay').click();
  assert.equal(element('nav').classList.contains('open'), false);
  element('menu').click();
  listeners.get('keydown')({ key: 'Escape' });
  assert.equal(element('nav').classList.contains('open'), false);
  assert.equal(element('menu').getAttribute('aria-expanded'), 'false');
});

test('Study, Projects and Models switch correct tab panels', () => {
  const { element } = harness();
  element('tab-projects').click();
  assert.equal(element('panel-study').hidden, true);
  assert.equal(element('panel-projects').hidden, false);
  assert.equal(element('panel-models').hidden, true);
  assert.equal(element('tab-projects').getAttribute('aria-selected'), 'true');
  element('tab-models').click();
  assert.equal(element('panel-models').hidden, false);
  assert.equal(element('panel-projects').hidden, true);
  element('tab-study').click();
  assert.equal(element('panel-study').hidden, false);
});

test('embedded browser only selects hardcoded Academy lessons', () => {
  const { element } = harness();
  element('lesson-agents').click();
  assert.equal(element('lesson-frame').getAttribute('src'), '../learn/ai-agents-101/index.html');
  element('lesson-evaluation').click();
  assert.equal(element('lesson-frame').getAttribute('src'), '../learn/evaluation-101/index.html');
  element('lesson-untrusted').click();
  assert.equal(element('lesson-frame').getAttribute('src'), '../learn/evaluation-101/index.html');
});

test('offline workspace makes no external AI requests or credential persistence', () => {
  assert.match(html, /id="send" type="button" disabled/);
  assert.match(html, /sandbox="allow-scripts"/);
  assert.doesNotMatch(script, /fetch\s*\(/);
  assert.doesNotMatch(script, /localStorage|sessionStorage/);
  assert.doesNotMatch(script, /token\.json/);
});
