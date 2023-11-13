"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[2672],{3905:(e,t,n)=>{n.d(t,{Zo:()=>p,kt:()=>g});var i=n(67294);function l(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function a(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);t&&(i=i.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,i)}return n}function o(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?a(Object(n),!0).forEach((function(t){l(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):a(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function r(e,t){if(null==e)return{};var n,i,l=function(e,t){if(null==e)return{};var n,i,l={},a=Object.keys(e);for(i=0;i<a.length;i++)n=a[i],t.indexOf(n)>=0||(l[n]=e[n]);return l}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(i=0;i<a.length;i++)n=a[i],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(l[n]=e[n])}return l}var c=i.createContext({}),s=function(e){var t=i.useContext(c),n=t;return e&&(n="function"==typeof e?e(t):o(o({},t),e)),n},p=function(e){var t=s(e.components);return i.createElement(c.Provider,{value:t},e.children)},h="mdxType",d={inlineCode:"code",wrapper:function(e){var t=e.children;return i.createElement(i.Fragment,{},t)}},u=i.forwardRef((function(e,t){var n=e.components,l=e.mdxType,a=e.originalType,c=e.parentName,p=r(e,["components","mdxType","originalType","parentName"]),h=s(n),u=l,g=h["".concat(c,".").concat(u)]||h[u]||d[u]||a;return n?i.createElement(g,o(o({ref:t},p),{},{components:n})):i.createElement(g,o({ref:t},p))}));function g(e,t){var n=arguments,l=t&&t.mdxType;if("string"==typeof e||l){var a=n.length,o=new Array(a);o[0]=u;var r={};for(var c in t)hasOwnProperty.call(t,c)&&(r[c]=t[c]);r.originalType=e,r[h]="string"==typeof e?e:l,o[1]=r;for(var s=2;s<a;s++)o[s]=n[s];return i.createElement.apply(null,o)}return i.createElement.apply(null,n)}u.displayName="MDXCreateElement"},75945:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>c,contentTitle:()=>o,default:()=>d,frontMatter:()=>a,metadata:()=>r,toc:()=>s});var i=n(87462),l=(n(67294),n(3905));const a={title:"ClientState",sidebar_label:"ClientState",sidebar_position:3,slug:"/ibc/light-clients/localhost/client-state"},o="ClientState",r={unversionedId:"light-clients/localhost/client-state",id:"light-clients/localhost/client-state",title:"ClientState",description:"The 09-localhost ClientState maintains a single field used to track the latest sequence of the state machine i.e. the height of the blockchain.",source:"@site/docs/03-light-clients/02-localhost/03-client-state.md",sourceDirName:"03-light-clients/02-localhost",slug:"/ibc/light-clients/localhost/client-state",permalink:"/main/ibc/light-clients/localhost/client-state",draft:!1,tags:[],version:"current",sidebarPosition:3,frontMatter:{title:"ClientState",sidebar_label:"ClientState",sidebar_position:3,slug:"/ibc/light-clients/localhost/client-state"},sidebar:"defaultSidebar",previous:{title:"Integration",permalink:"/main/ibc/light-clients/localhost/integration"},next:{title:"Connection",permalink:"/main/ibc/light-clients/localhost/connection"}},c={},s=[{value:"Client updates",id:"client-updates",level:2}],p={toc:s},h="wrapper";function d(e){let{components:t,...n}=e;return(0,l.kt)(h,(0,i.Z)({},p,n,{components:t,mdxType:"MDXLayout"}),(0,l.kt)("h1",{id:"clientstate"},(0,l.kt)("inlineCode",{parentName:"h1"},"ClientState")),(0,l.kt)("p",null,"The 09-localhost ",(0,l.kt)("inlineCode",{parentName:"p"},"ClientState")," maintains a single field used to track the latest sequence of the state machine i.e. the height of the blockchain."),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-go"},"type ClientState struct {\n  // the latest height of the blockchain\n  LatestHeight clienttypes.Height\n}\n")),(0,l.kt)("p",null,"The 09-localhost ",(0,l.kt)("inlineCode",{parentName:"p"},"ClientState")," is instantiated in the ",(0,l.kt)("inlineCode",{parentName:"p"},"InitGenesis")," handler of the 02-client submodule in core IBC.\nIt calls ",(0,l.kt)("inlineCode",{parentName:"p"},"CreateLocalhostClient"),", declaring a new ",(0,l.kt)("inlineCode",{parentName:"p"},"ClientState")," and initializing it with its own client prefixed store."),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-go"},"func (k Keeper) CreateLocalhostClient(ctx sdk.Context) error {\n  var clientState localhost.ClientState\n  return clientState.Initialize(ctx, k.cdc, k.ClientStore(ctx, exported.LocalhostClientID), nil)\n}\n")),(0,l.kt)("p",null,"It is possible to disable the localhost client by removing the ",(0,l.kt)("inlineCode",{parentName:"p"},"09-localhost")," entry from the ",(0,l.kt)("inlineCode",{parentName:"p"},"allowed_clients")," list through governance."),(0,l.kt)("h2",{id:"client-updates"},"Client updates"),(0,l.kt)("p",null,"The latest height is updated periodically through the ABCI ",(0,l.kt)("a",{parentName:"p",href:"https://docs.cosmos.network/v0.47/building-modules/beginblock-endblock"},(0,l.kt)("inlineCode",{parentName:"a"},"BeginBlock"))," interface of the 02-client submodule in core IBC."),(0,l.kt)("p",null,(0,l.kt)("a",{parentName:"p",href:"https://github.com/cosmos/ibc-go/blob/09-localhost/modules/core/02-client/abci.go#L12"},"See ",(0,l.kt)("inlineCode",{parentName:"a"},"BeginBlocker")," in abci.go.")),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-go"},"func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {\n  // ...\n\n  if clientState, found := k.GetClientState(ctx, exported.Localhost); found {\n    if k.GetClientStatus(ctx, clientState, exported.Localhost) == exported.Active {\n      k.UpdateLocalhostClient(ctx, clientState)\n    }\n  }\n}\n")),(0,l.kt)("p",null,"The above calls into the 09-localhost ",(0,l.kt)("inlineCode",{parentName:"p"},"UpdateState")," method of the ",(0,l.kt)("inlineCode",{parentName:"p"},"ClientState")," .\nIt retrieves the current block height from the application context and sets the ",(0,l.kt)("inlineCode",{parentName:"p"},"LatestHeight")," of the 09-localhost client."),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-go"},"func (cs ClientState) UpdateState(ctx sdk.Context, cdc codec.BinaryCodec, clientStore sdk.KVStore, clientMsg exported.ClientMessage) []exported.Height {\n  height := clienttypes.GetSelfHeight(ctx)\n  cs.LatestHeight = height\n\n  clientStore.Set(host.ClientStateKey(), clienttypes.MustMarshalClientState(cdc, &cs))\n\n  return []exported.Height{height}\n}\n")),(0,l.kt)("p",null,"Note that the 09-localhost ",(0,l.kt)("inlineCode",{parentName:"p"},"ClientState")," is not updated through the 02-client interface leveraged by conventional IBC light clients."))}d.isMDXComponent=!0}}]);