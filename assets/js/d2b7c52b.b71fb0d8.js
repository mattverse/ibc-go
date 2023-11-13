"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[3768],{3905:(e,t,n)=>{n.d(t,{Zo:()=>p,kt:()=>g});var i=n(67294);function o(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function r(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);t&&(i=i.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,i)}return n}function a(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?r(Object(n),!0).forEach((function(t){o(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):r(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function l(e,t){if(null==e)return{};var n,i,o=function(e,t){if(null==e)return{};var n,i,o={},r=Object.keys(e);for(i=0;i<r.length;i++)n=r[i],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);for(i=0;i<r.length;i++)n=r[i],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}var s=i.createContext({}),c=function(e){var t=i.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):a(a({},t),e)),n},p=function(e){var t=c(e.components);return i.createElement(s.Provider,{value:t},e.children)},u="mdxType",d={inlineCode:"code",wrapper:function(e){var t=e.children;return i.createElement(i.Fragment,{},t)}},m=i.forwardRef((function(e,t){var n=e.components,o=e.mdxType,r=e.originalType,s=e.parentName,p=l(e,["components","mdxType","originalType","parentName"]),u=c(n),m=o,g=u["".concat(s,".").concat(m)]||u[m]||d[m]||r;return n?i.createElement(g,a(a({ref:t},p),{},{components:n})):i.createElement(g,a({ref:t},p))}));function g(e,t){var n=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var r=n.length,a=new Array(r);a[0]=m;var l={};for(var s in t)hasOwnProperty.call(t,s)&&(l[s]=t[s]);l.originalType=e,l[u]="string"==typeof e?e:o,a[1]=l;for(var c=2;c<r;c++)a[c]=n[c];return i.createElement.apply(null,a)}return i.createElement.apply(null,n)}m.displayName="MDXCreateElement"},24906:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>s,contentTitle:()=>a,default:()=>d,frontMatter:()=>r,metadata:()=>l,toc:()=>c});var i=n(87462),o=(n(67294),n(3905));const r={title:"Setup",sidebar_label:"Setup",sidebar_position:9,slug:"/ibc/light-clients/setup"},a="Setup",l={unversionedId:"light-clients/developer-guide/setup",id:"version-v7.3.x/light-clients/developer-guide/setup",title:"Setup",description:"Learn how to configure light client modules and create clients using core IBC and the 02-client submodule.",source:"@site/versioned_docs/version-v7.3.x/03-light-clients/01-developer-guide/09-setup.md",sourceDirName:"03-light-clients/01-developer-guide",slug:"/ibc/light-clients/setup",permalink:"/v7/ibc/light-clients/setup",draft:!1,tags:[],version:"v7.3.x",sidebarPosition:9,frontMatter:{title:"Setup",sidebar_label:"Setup",sidebar_position:9,slug:"/ibc/light-clients/setup"},sidebar:"defaultSidebar",previous:{title:"Handling Genesis",permalink:"/v7/ibc/light-clients/genesis"},next:{title:"Solomachine",permalink:"/v7/ibc/light-clients/solomachine/solomachine"}},s={},c=[{value:"Configuring a light client module",id:"configuring-a-light-client-module",level:2},{value:"Creating clients",id:"creating-clients",level:2}],p={toc:c},u="wrapper";function d(e){let{components:t,...n}=e;return(0,o.kt)(u,(0,i.Z)({},p,n,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("h1",{id:"setup"},"Setup"),(0,o.kt)("admonition",{title:"Synopsis",type:"note"},(0,o.kt)("p",{parentName:"admonition"},"Learn how to configure light client modules and create clients using core IBC and the ",(0,o.kt)("inlineCode",{parentName:"p"},"02-client")," submodule. ")),(0,o.kt)("p",null,"A last step to finish the development of the light client, is to implement the ",(0,o.kt)("inlineCode",{parentName:"p"},"AppModuleBasic")," interface to allow it to be added to the chain's ",(0,o.kt)("inlineCode",{parentName:"p"},"app.go")," alongside other light client types the chain enables."),(0,o.kt)("p",null,"Finally, a succinct rundown is given of the remaining steps to make the light client operational, getting the light client type passed through governance and creating the clients."),(0,o.kt)("h2",{id:"configuring-a-light-client-module"},"Configuring a light client module"),(0,o.kt)("p",null,"An IBC light client module must implement the ",(0,o.kt)("a",{parentName:"p",href:"https://github.com/cosmos/cosmos-sdk/blob/main/types/module/module.go#L50"},(0,o.kt)("inlineCode",{parentName:"a"},"AppModuleBasic"))," interface in order to register its concrete types against the core IBC interfaces defined in ",(0,o.kt)("inlineCode",{parentName:"p"},"modules/core/exported"),". This is accomplished via the ",(0,o.kt)("inlineCode",{parentName:"p"},"RegisterInterfaces")," method which provides the light client module with the opportunity to register codec types using the chain's ",(0,o.kt)("inlineCode",{parentName:"p"},"InterfaceRegistry"),". Please refer to the ",(0,o.kt)("a",{parentName:"p",href:"https://github.com/cosmos/ibc-go/blob/02-client-refactor-beta1/modules/light-clients/07-tendermint/codec.go#L11"},(0,o.kt)("inlineCode",{parentName:"a"},"07-tendermint")," codec registration"),"."),(0,o.kt)("p",null,"The ",(0,o.kt)("inlineCode",{parentName:"p"},"AppModuleBasic")," interface may also be leveraged to install custom CLI handlers for light client module users. Light client modules can safely no-op for interface methods which it does not wish to implement."),(0,o.kt)("p",null,"Please refer to the ",(0,o.kt)("a",{parentName:"p",href:"/v7/ibc/integration#integrating-light-clients"},"core IBC documentation")," for how to configure additional light client modules alongside ",(0,o.kt)("inlineCode",{parentName:"p"},"07-tendermint")," in ",(0,o.kt)("inlineCode",{parentName:"p"},"app.go"),"."),(0,o.kt)("p",null,"See below for an example of the ",(0,o.kt)("inlineCode",{parentName:"p"},"07-tendermint")," implementation of ",(0,o.kt)("inlineCode",{parentName:"p"},"AppModuleBasic"),"."),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},"var _ module.AppModuleBasic = AppModuleBasic{}\n\n// AppModuleBasic defines the basic application module used by the tendermint light client.\n// Only the RegisterInterfaces function needs to be implemented. All other function perform\n// a no-op.\ntype AppModuleBasic struct{}\n\n// Name returns the tendermint module name.\nfunc (AppModuleBasic) Name() string {\n  return ModuleName\n}\n\n// RegisterLegacyAminoCodec performs a no-op. The Tendermint client does not support amino.\nfunc (AppModuleBasic) RegisterLegacyAminoCodec(*codec.LegacyAmino) {}\n\n// RegisterInterfaces registers module concrete types into protobuf Any. This allows core IBC\n// to unmarshal tendermint light client types.\nfunc (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {\n  RegisterInterfaces(registry)\n}\n\n// DefaultGenesis performs a no-op. Genesis is not supported for the tendermint light client.\nfunc (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {\n  return nil\n}\n\n// ValidateGenesis performs a no-op. Genesis is not supported for the tendermint light cilent.\nfunc (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {\n  return nil\n}\n\n// RegisterGRPCGatewayRoutes performs a no-op.\nfunc (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {}\n\n// GetTxCmd performs a no-op. Please see the 02-client cli commands.\nfunc (AppModuleBasic) GetTxCmd() *cobra.Command {\n  return nil\n}\n\n// GetQueryCmd performs a no-op. Please see the 02-client cli commands.\nfunc (AppModuleBasic) GetQueryCmd() *cobra.Command {\n  return nil\n}\n")),(0,o.kt)("h2",{id:"creating-clients"},"Creating clients"),(0,o.kt)("p",null,"A client is created by executing a new ",(0,o.kt)("inlineCode",{parentName:"p"},"MsgCreateClient")," transaction composed with a valid ",(0,o.kt)("inlineCode",{parentName:"p"},"ClientState")," and initial ",(0,o.kt)("inlineCode",{parentName:"p"},"ConsensusState")," encoded as protobuf ",(0,o.kt)("inlineCode",{parentName:"p"},"Any"),"s.\nGenerally, this is performed by an off-chain process known as an ",(0,o.kt)("a",{parentName:"p",href:"https://github.com/cosmos/ibc/tree/main/spec/relayer/ics-018-relayer-algorithms"},"IBC relayer")," however, this is not a strict requirement."),(0,o.kt)("p",null,"See below for a list of IBC relayer implementations:"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"https://github.com/cosmos/relayer"},"cosmos/relayer")),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"https://github.com/informalsystems/hermes"},"informalsystems/hermes")),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"https://github.com/confio/ts-relayer"},"confio/ts-relayer"))),(0,o.kt)("p",null,"Stateless checks are performed within the ",(0,o.kt)("a",{parentName:"p",href:"https://github.com/cosmos/ibc-go/blob/02-client-refactor-beta1/modules/core/02-client/types/msgs.go#L48"},(0,o.kt)("inlineCode",{parentName:"a"},"ValidateBasic"))," method of ",(0,o.kt)("inlineCode",{parentName:"p"},"MsgCreateClient"),"."),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-protobuf"},'// MsgCreateClient defines a message to create an IBC client\nmessage MsgCreateClient {\n  option (gogoproto.equal)           = false;\n  option (gogoproto.goproto_getters) = false;\n\n  // light client state\n  google.protobuf.Any client_state = 1 [(gogoproto.moretags) = "yaml:\\"client_state\\""];\n  // consensus state associated with the client that corresponds to a given\n  // height.\n  google.protobuf.Any consensus_state = 2 [(gogoproto.moretags) = "yaml:\\"consensus_state\\""];\n  // signer address\n  string signer = 3;\n}\n')),(0,o.kt)("p",null,"Leveraging protobuf ",(0,o.kt)("inlineCode",{parentName:"p"},"Any")," encoding allows core IBC to ",(0,o.kt)("a",{parentName:"p",href:"https://github.com/cosmos/ibc-go/blob/02-client-refactor-beta1/modules/core/keeper/msg_server.go#L28-L36"},"unpack")," both the ",(0,o.kt)("inlineCode",{parentName:"p"},"ClientState")," and ",(0,o.kt)("inlineCode",{parentName:"p"},"ConsensusState")," into their respective interface types registered previously using the light client module's ",(0,o.kt)("inlineCode",{parentName:"p"},"RegisterInterfaces")," method."),(0,o.kt)("p",null,"Within the ",(0,o.kt)("inlineCode",{parentName:"p"},"02-client")," submodule, the ",(0,o.kt)("a",{parentName:"p",href:"https://github.com/cosmos/ibc-go/blob/02-client-refactor-beta1/modules/core/02-client/keeper/client.go#L30-L34"},(0,o.kt)("inlineCode",{parentName:"a"},"ClientState")," is then initialized")," with its own isolated key-value store, namespaced using a unique client identifier."),(0,o.kt)("p",null,"In order to successfully create an IBC client using a new client type, it ",(0,o.kt)("a",{parentName:"p",href:"https://github.com/cosmos/ibc-go/blob/02-client-refactor-beta1/modules/core/02-client/keeper/client.go#L18-L24"},"must be supported"),". Light client support in IBC is gated by on-chain governance. The allow list may be updated by submitting a new governance proposal to update the ",(0,o.kt)("inlineCode",{parentName:"p"},"02-client")," parameter ",(0,o.kt)("inlineCode",{parentName:"p"},"AllowedClients"),"."),(0,o.kt)("p",null,"See below for example:"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-shell"},"%s tx gov submit-proposal param-change <path/to/proposal.json> --from=<key_or_address>\n")),(0,o.kt)("p",null,"where ",(0,o.kt)("inlineCode",{parentName:"p"},"proposal.json")," contains:"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-json"},'{\n  "title": "IBC Clients Param Change",\n  "description": "Update allowed clients",\n  "changes": [\n    {\n      "subspace": "ibc",\n      "key": "AllowedClients",\n      "value": ["06-solomachine", "07-tendermint", "0x-new-client"]\n    }\n  ],\n  "deposit": "1000stake"\n}\n')))}d.isMDXComponent=!0}}]);