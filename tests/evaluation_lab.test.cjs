// Dependency-free behavior tests for the original browser-only evaluation lab.
const test=require('node:test');
const assert=require('node:assert/strict');
const fs=require('node:fs');
const path=require('node:path');
const vm=require('node:vm');
const source=fs.readFileSync(path.join(__dirname,'..','site','lab','evaluation','index.html'),'utf8');
const script=source.match(/<script>([\s\S]*?)<\/script>/)?.[1];
assert.ok(script,'lab JavaScript not found');
function boot(){
 const defaults={tp:'70',fp:'10',fn:'20',tn:'100'};
 const elements={};
 for(const id of ['tp','fp','fn','tn','precision','recall','specificity','f1','accuracy','support','matrix-tp','matrix-fp','matrix-fn','matrix-tn','message','reset','export']){
  elements[id]={value:defaults[id]??'',textContent:'',disabled:false,handlers:{},addEventListener(name,fn){this.handlers[name]=fn;}};
 }
 const downloads=[];
 const document={getElementById:id=>elements[id],createElement:()=>({set href(v){this._href=v},set download(v){this._name=v},click(){downloads.push({href:this._href,name:this._name})}})};
 const URL={createObjectURL:()=> 'blob:practice',revokeObjectURL:()=>{}};
 const Blob=class{constructor(parts,opts){this.parts=parts;this.opts=opts}};
 vm.runInNewContext(script,{document,URL,Blob,setTimeout:fn=>fn()});
 return {e:elements,downloads};
}
test('default confusion matrix metrics',()=>{
 const {e}=boot();
 assert.equal(e.precision.textContent,'87.5%');
 assert.equal(e.recall.textContent,'77.8%');
 assert.equal(e.specificity.textContent,'90.9%');
 assert.equal(e.f1.textContent,'82.4%');
 assert.equal(e.accuracy.textContent,'85.0%');
 assert.equal(e.support.textContent,'90');
});
test('undefined metrics use N/A rather than zero or NaN',()=>{
 const {e}=boot();
 for(const key of ['tp','fp','fn','tn'])e[key].value='0';
 e.tp.handlers.input();
 for(const key of ['precision','recall','specificity','f1','accuracy'])assert.equal(e[key].textContent,'N/A');
 assert.equal(e.support.textContent,'0');
});
test('invalid counts disable export and explain error',()=>{
 const {e}=boot();
 for(const bad of ['-1','1.5','1000001','']){
  e.tp.value=bad;e.tp.handlers.input();
  assert.equal(e.export.disabled,true,bad);
  assert.match(e.message.textContent,/whole number/);
 }
});
test('reset restores counts and recomputes',()=>{
 const {e}=boot();
 e.tp.value='0';e.tp.handlers.input();
 e.reset.handlers.click();
 assert.equal(e.tp.value,'70');
 assert.equal(e.precision.textContent,'87.5%');
});
test('export is a local JSON practice receipt, not a credential',()=>{
 const {e,downloads}=boot();
 e.export.handlers.click();
 assert.equal(downloads.length,1);
 assert.equal(downloads[0].name,'academy-evaluation-practice.json');
});
