(function(a,s){typeof exports=="object"&&typeof module<"u"?s(exports):typeof define=="function"&&define.amd?define(["exports"],s):(a=typeof globalThis<"u"?globalThis:a||self,s(a.XErr={}))})(this,function(a){"use strict";var v=Object.defineProperty;var L=(a,s,d)=>s in a?v(a,s,{enumerable:!0,configurable:!0,writable:!0,value:d}):a[s]=d;var r=(a,s,d)=>L(a,typeof s!="symbol"?s+"":s,d);const s=[];for(let t=0;t<256;++t)s.push((t+256).toString(16).slice(1));function d(t,e=0){return(s[t[e+0]]+s[t[e+1]]+s[t[e+2]]+s[t[e+3]]+"-"+s[t[e+4]]+s[t[e+5]]+"-"+s[t[e+6]]+s[t[e+7]]+"-"+s[t[e+8]]+s[t[e+9]]+"-"+s[t[e+10]]+s[t[e+11]]+s[t[e+12]]+s[t[e+13]]+s[t[e+14]]+s[t[e+15]]).toLowerCase()}let m;const g=new Uint8Array(16);function y(){if(!m){if(typeof crypto>"u"||!crypto.getRandomValues)throw new Error("crypto.getRandomValues() not supported. See https://github.com/uuidjs/uuid#getrandomvalues-not-supported");m=crypto.getRandomValues.bind(crypto)}return m(g)}const h={};function w(t,e,i){let n;{const l=Date.now(),o=y();_(h,l,o),n=S(o,h.msecs,h.nsecs,h.clockseq,h.node,e,i)}return e?n:d(n)}function _(t,e,i){return t.msecs??(t.msecs=-1/0),t.nsecs??(t.nsecs=0),e===t.msecs?(t.nsecs++,t.nsecs>=1e4&&(t.node=void 0,t.nsecs=0)):e>t.msecs?t.nsecs=0:e<t.msecs&&(t.node=void 0),t.node||(t.node=i.slice(10,16),t.node[0]|=1,t.clockseq=(i[8]<<8|i[9])&16383),t.msecs=e,t}function S(t,e,i,n,l,o,c=0){o||(o=new Uint8Array(16),c=0),e??(e=Date.now()),i??(i=0),n??(n=(t[8]<<8|t[9])&16383),l??(l=t.slice(10,16)),e+=122192928e5;const u=((e&268435455)*1e4+i)%4294967296;o[c++]=u>>>24&255,o[c++]=u>>>16&255,o[c++]=u>>>8&255,o[c++]=u&255;const p=e/4294967296*1e4&268435455;o[c++]=p>>>8&255,o[c++]=p&255,o[c++]=p>>>24&15|16,o[c++]=p>>>16&255,o[c++]=n>>>8|128,o[c++]=n&255;for(let f=0;f<6;++f)o[c++]=l[f];return o}class x{constructor(e,i){r(this,"Dns","");r(this,"client_id","");r(this,"Pid","");r(this,"Uid","");r(this,"platform",null);r(this,"MessageList",[]);r(this,"timer",0);r(this,"Push",e=>{var i;this.MessageList.push({...e,ProjectKey:this.Pid,ClientId:this.client_id}),this.MessageList.length>5?this.upload():(i=this.platform)==null||i.setCache("x_err_message_list",this.MessageList)});r(this,"uploadInfo",e=>{var i;if(this.Dns)try{(i=this.platform)==null||i.upload(this.Dns+"/admin/monitor_client/add",{ProjectKey:this.Pid,ClientId:this.client_id,UserId:this.Uid,Width:e.ScreenWidth,Height:e.ScreenHeight}).catch(n=>{})}catch{}});r(this,"uploadSlow",e=>{var i;if(this.Dns)try{(i=this.platform)==null||i.upload(this.Dns+"/admin/monitor_slow/add",{ProjectKey:this.Pid,ClientId:this.client_id,UserId:this.Uid,Path:e.Path,Time:e.Time}).catch(n=>{})}catch{}});if(!e){console.error("props is null");return}if(!i){console.error("platform is null");return}if(this.platform=i,e.Dns&&e.Pid)this.Dns=e.Dns,this.Pid=e.Pid;else{console.error("props.Dns and props.Pid cannot be null");return}e.Uid&&(this.Uid=String(e.Uid)),this.setClientID(),this.SetUid(),this.getLocalMessage(),i.listen(n=>{if(console.log("listenCallback",n),n.Type=="onloadTime"){let l=n;this.uploadSlow(l)}else this.Push(n)}),this.timer=setInterval(()=>{this.upload()},1e3*10)}SetUid(e){var i,n;if(e)this.Uid=String(e),(i=this.platform)==null||i.setCache("x_err_uid",this.Uid);else{const l=(n=this.platform)==null?void 0:n.getCache("x_err_uid");l&&(this.Uid=l)}this.initEnv()}initEnv(){var i;let e=(i=this.platform)==null?void 0:i.getEnvInfo();e&&this.uploadInfo(e)}setClientID(){var i,n;const e=(i=this.platform)==null?void 0:i.getCache("x_err_client_id");e?this.client_id=e:(this.client_id=w(),(n=this.platform)==null||n.setCache("x_err_client_id",this.client_id))}getLocalMessage(){var i;let e=(i=this.platform)==null?void 0:i.getCache("x_err_message_list");e?this.MessageList=e:this.MessageList=[]}upload(){var e,i;if(this.Dns&&this.MessageList.length)try{(e=this.platform)==null||e.upload(this.Dns+"/admin/monitor_error/add",this.MessageList).catch(n=>{}),this.MessageList=[],(i=this.platform)==null||i.delCache("x_err_message_list")}catch{}}unListen(){var e;clearInterval(this.timer),(e=this.platform)==null||e.unListen()}}class T{constructor(e){r(this,"props");r(this,"listenError",e=>{var n;console.error([e]);let i=e.target;i!=null&&i.localName?(i==null?void 0:i.localName)==="img"||(i==null?void 0:i.localName)==="script"?this.callback({Type:"resources",EventType:i==null?void 0:i.localName,Path:i.src,Message:"",Stack:""}):(i==null?void 0:i.localName)==="link"&&this.callback({Type:"resources",EventType:i==null?void 0:i.localName,Path:i.href}):this.callback({Type:"error",EventType:e.type,Path:window.location.href,Message:e.message,Stack:this.handleStack(((n=e.error)==null?void 0:n.stack)||"")})});r(this,"unhandledrejection",e=>{var i,n;console.error(e),e&&typeof e.reason=="string"?this.callback({Type:"error",EventType:e.type,Path:window.location.href,Message:e.reason,Stack:""}):e&&typeof e.reason=="object"&&this.callback({Type:"error",EventType:e.type,Path:window.location.href,Message:((i=e.reason)==null?void 0:i.message)||"",Stack:this.handleStack(((n=e.reason)==null?void 0:n.stack)||"")})});r(this,"onLoad",()=>{const e=performance.getEntriesByType("navigation");if(e.length>0){const i=e[0];console.log("performanceData",i);let n=i.loadEventStart-i.startTime;this.props.onloadTimeOut&&n>this.props.onloadTimeOut&&this.callback({Type:"onloadTime",Path:window.location.href,Time:n})}});this.props={onloadTimeOut:5e3,...e}}upload(e,i){return new Promise(n=>{try{let l=new Image;l.onload=function(o){var c=o;c.preventDefault()},l.onerror=function(){},l.src=e+"?data="+encodeURIComponent(JSON.stringify(i)),n()}catch{n()}})}setCache(e,i){localStorage.setItem(e,JSON.stringify(i))}getCache(e){try{let i=localStorage.getItem(e);return i?JSON.parse(i):null}catch{return null}}delCache(e){localStorage.removeItem(e)}getEnvInfo(){const e={Type:"env",ScreenHeight:0,ScreenWidth:0};return window&&(e.ScreenHeight=window.innerHeight||0,e.ScreenWidth=window.innerWidth||0),e}callback(e){}handleStack(e){let i=[];return e&&e.split(`
`).map((n,l)=>{l<4&&i.push(n)}),i.join(`
`)}listen(e){this.callback=e,window.addEventListener("unhandledrejection",this.unhandledrejection),window.addEventListener("error",this.listenError,!0),window.addEventListener("load",this.onLoad)}unListen(){this.callback=()=>{},window.removeEventListener("error",this.listenError),window.removeEventListener("unhandledrejection",this.unhandledrejection)}}a.XErr=x,a.XErrWeb=T,Object.defineProperty(a,Symbol.toStringTag,{value:"Module"})});
