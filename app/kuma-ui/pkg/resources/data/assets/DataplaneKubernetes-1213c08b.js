import{d as Y,r as m,c as N,q as Q,o as u,h as p,g as e,e as i,w as t,f as a,t as h,u as r,W as c,a6 as $,F as H,l as X,a7 as _,a as x,a5 as M,b as y,p as J,k as Z}from"./runtime-dom.esm-bundler-9284044f.js";import{M as v,a as C,b as B}from"./kongponents.es-3ba46133.js";import{a as ee,z as ae,P as se}from"./production-554ae9d4.js";import{F as k,E as ne,S as te}from"./StepSkeleton-6deb7529.js";import{E as le}from"./EnvironmentSwitcher-c84f3619.js";import{f as oe}from"./formatForCLI-931cd5c6.js";import{_ as P}from"./CodeBlock.vue_vue_type_style_index_0_lang-c7bd4e33.js";import{u as ie}from"./store-f2021894.js";import{u as re}from"./index-36b3783c.js";import{Q as de}from"./QueryParameter-70743f73.js";import{_ as ue}from"./_plugin-vue_export-helper-c27b6911.js";import"./toYaml-4e00099e.js";const ce={apiVersion:"v1",kind:"Namespace",metadata:{name:null,namespace:null,annotations:{"kuma.io/sidecar-injection":"enabled","kuma.io/mesh":null}}},l=f=>(J("data-v-f14bf435"),f=f(),Z(),f),pe={class:"wizard"},me={class:"wizard__content"},he=l(()=>e("h3",null,`
            Create Kubernetes Dataplane
          `,-1)),ve=l(()=>e("h3",null,`
            To get started, please select on what Mesh you would like to add the Dataplane:
          `,-1)),ke=l(()=>e("p",null,`
            If you've got an existing Mesh that you would like to associate with your
            Dataplane, you can select it below, or create a new one using our Mesh Wizard.
          `,-1)),_e=l(()=>e("small",null,"Would you like to see instructions for Universal? Use sidebar to change wizard!",-1)),ye=l(()=>e("option",{disabled:"",value:""},`
                      Select an existing Mesh…
                    `,-1)),ge=["value"],fe=l(()=>e("label",{class:"k-input-label mr-4"},`
                    or
                  `,-1)),be=l(()=>e("h3",null,`
            Setup Dataplane Mode
          `,-1)),we=l(()=>e("p",null,`
            You can create a data plane for a service or a data plane for a Gateway.
          `,-1)),Se={for:"service-dataplane"},De=l(()=>e("span",null,`
                    Service Dataplane
                  `,-1)),Ie={for:"ingress-dataplane"},Ne=l(()=>e("span",null,`
                    Ingress Dataplane
                  `,-1)),xe={key:0},Me=l(()=>e("p",null,`
              Should the data plane be added for an entire Namespace and all of its services,
              or for specific individual services in any namespace?
            `,-1)),Ee={for:"k8s-services-all"},Te=l(()=>e("span",null,`
                      All Services in Namespace
                    `,-1)),Ve={for:"k8s-services-individual"},Ce=l(()=>e("span",null,`
                      Individual Services
                    `,-1)),Be={key:1},Pe={for:"k8s-ingress-kong"},Ue=l(()=>e("span",null,`
                      Kong Ingress
                    `,-1)),Ke={for:"k8s-ingress-other"},Fe=l(()=>e("span",null,`
                      Other Ingress
                    `,-1)),Ae=l(()=>e("p",null,`
                  Please go ahead and deploy the Ingress first, then restart this wizard and select “Existing Ingress”.
                `,-1)),qe={key:0},ze={key:0},We=l(()=>e("h3",null,`
                Auto-Inject DPP
              `,-1)),je=l(()=>e("p",null,`
                You can now execute the following commands to automatically inject the sidecar proxy in every Pod, and by doing so creating the Dataplane.
              `,-1)),Oe=l(()=>e("h4",null,"Kubernetes",-1)),Ge=l(()=>e("h3",null,"Searching…",-1)),Le=l(()=>e("p",null,"We are looking for your dataplane.",-1)),Re=l(()=>e("h3",null,"Done!",-1)),Ye={key:0},Qe=l(()=>e("p",null,`
                  Proceed to the next step where we will show you
                  your new Dataplane.
                `,-1)),$e=l(()=>e("h3",null,"Mesh not found",-1)),He=l(()=>e("p",null,"We were unable to find your mesh.",-1)),Xe=l(()=>e("p",null,`
                Please return to the first step and make sure to select an
                existing Mesh, or create a new one.
              `,-1)),Je=l(()=>e("h3",null,"Dataplane",-1)),Ze=l(()=>e("h3",null,"Example",-1)),ea=l(()=>e("p",null,`
            Below is an example of a Dataplane resource output:
          `,-1)),aa=Y({__name:"DataplaneKubernetes",setup(f){const U=re(),K=`apiVersion: 'kuma.io/v1alpha1'
kind: Dataplane
mesh: default
metadata:
  name: dp-echo-1
  annotations:
    kuma.io/sidecar-injection: enabled
    kuma.io/mesh: default
networking:
  address: 10.0.0.1
  inbound:
  - port: 10000
    servicePort: 9000
    tags:
      kuma.io/service: echo`,F=[{label:"General",slug:"general"},{label:"Scope Settings",slug:"scope-settings"},{label:"Install",slug:"complete"}],A=[{name:"dataplane"},{name:"example"},{name:"switch"}],q=ee(),S=ie(),z=m(ce),D=m(0),E=m(!1),I=m(!1),b=m(!1),T=m(!1),s=m({meshName:"",k8sDataplaneType:"dataplane-type-service",k8sServices:"all-services",k8sNamespace:"",k8sNamespaceSelection:"",k8sServiceDeployment:"",k8sServiceDeploymentSelection:"",k8sIngressDeployment:"",k8sIngressDeploymentSelection:"",k8sIngressType:"",k8sIngressBrand:"kong-ingress",k8sIngressSelection:""}),w=N(()=>S.getters["config/getTagline"]),W=N(()=>{const d=Object.assign({},z.value),n=s.value.k8sNamespaceSelection;if(!n)return"";d.metadata.name=n,d.metadata.namespace=n,d.metadata.annotations["kuma.io/mesh"]=s.value.meshName;const o=`" | kubectl apply -f - && kubectl delete pod --all -n ${n}`;return oe(d,o)}),j=N(()=>{const{k8sNamespaceSelection:d,meshName:n}=s.value;return n.length===0?!0:D.value===1?!d:!1});Q(()=>s.value.k8sNamespaceSelection,function(d){s.value.k8sNamespaceSelection=ae(d)});const V=de.get("step");D.value=V!==null?parseInt(V):0;function O(d){D.value=d}function G(){I.value=!0}async function L(){const n=s.value.meshName,o=s.value.k8sNamespaceSelection;if(T.value=!1,b.value=!1,!(!n||!o))try{const g=await U.getDataplaneFromMesh({mesh:n,name:o});g&&g.name.length>0?E.value=!0:b.value=!0}catch(g){b.value=!0,console.error(g)}finally{T.value=!0}}function R(){S.dispatch("updateSelectedMesh",s.value.meshName),q.push({name:"data-plane-list-view",params:{mesh:s.value.meshName}})}return(d,n)=>(u(),p("div",pe,[e("div",me,[i(te,{steps:F,"sidebar-content":A,"footer-enabled":I.value===!1,"next-disabled":r(j),onGoToStep:O},{general:t(()=>[he,a(),e("p",null,`
            Welcome to the wizard to create a new Dataplane resource in `+h(r(w))+`.
            We will be providing you with a few steps that will get you started.
          `,1),a(),e("p",null,`
            As you know, the `+h(r(se))+` GUI is read-only.
          `,1),a(),ve,a(),ke,a(),_e,a(),i(r(v),{class:"my-6","has-shadow":""},{body:t(()=>[i(k,{title:"Choose a Mesh","for-attr":"dp-mesh","all-inline":""},{default:t(()=>[e("div",null,[c(e("select",{id:"dp-mesh","onUpdate:modelValue":n[0]||(n[0]=o=>s.value.meshName=o),class:"k-input w-100"},[ye,a(),(u(!0),p(H,null,X(r(S).getters.getMeshList.items,o=>(u(),p("option",{key:o.name,value:o.name},h(o.name),9,ge))),128))],512),[[$,s.value.meshName]])]),a(),e("div",null,[fe,a(),i(r(C),{to:{name:"create-mesh"},appearance:"outline"},{default:t(()=>[a(`
                    Create a new Mesh
                  `)]),_:1})])]),_:1})]),_:1})]),"scope-settings":t(()=>[be,a(),we,a(),i(r(v),{class:"my-6","has-shadow":""},{body:t(()=>[i(k,{"all-inline":"","equal-cols":"","hide-label-col":""},{default:t(()=>[e("label",Se,[c(e("input",{id:"service-dataplane","onUpdate:modelValue":n[1]||(n[1]=o=>s.value.k8sDataplaneType=o),class:"k-input",type:"radio",name:"dataplane-type",value:"dataplane-type-service",checked:""},null,512),[[_,s.value.k8sDataplaneType]]),a(),De]),a(),e("label",Ie,[c(e("input",{id:"ingress-dataplane","onUpdate:modelValue":n[2]||(n[2]=o=>s.value.k8sDataplaneType=o),class:"k-input",type:"radio",name:"dataplane-type",value:"dataplane-type-ingress",disabled:""},null,512),[[_,s.value.k8sDataplaneType]]),a(),Ne])]),_:1})]),_:1}),a(),s.value.k8sDataplaneType==="dataplane-type-service"?(u(),p("div",xe,[Me,a(),i(r(v),{class:"my-6","has-shadow":""},{body:t(()=>[i(k,{"all-inline":"","equal-cols":"","hide-label-col":""},{default:t(()=>[e("label",Ee,[c(e("input",{id:"k8s-services-all","onUpdate:modelValue":n[3]||(n[3]=o=>s.value.k8sServices=o),class:"k-input",type:"radio",name:"k8s-services",value:"all-services",checked:""},null,512),[[_,s.value.k8sServices]]),a(),Te]),a(),e("label",Ve,[c(e("input",{id:"k8s-services-individual","onUpdate:modelValue":n[4]||(n[4]=o=>s.value.k8sServices=o),class:"k-input",type:"radio",name:"k8s-services",value:"individual-services",disabled:""},null,512),[[_,s.value.k8sServices]]),a(),Ce])]),_:1})]),_:1}),a(),s.value.k8sServices==="individual-services"?(u(),x(r(v),{key:0,class:"my-6","has-shadow":""},{body:t(()=>[i(k,{title:"Deployments","for-attr":"k8s-deployment-selection"},{default:t(()=>[c(e("input",{id:"k8s-service-deployment-new","onUpdate:modelValue":n[5]||(n[5]=o=>s.value.k8sServiceDeploymentSelection=o),type:"text",class:"k-input w-100",placeholder:"your-new-deployment",required:""},null,512),[[M,s.value.k8sServiceDeploymentSelection]])]),_:1})]),_:1})):y("",!0),a(),i(r(v),{class:"my-6","has-shadow":""},{body:t(()=>[i(k,{title:"Namespace","for-attr":"k8s-namespace-selection"},{default:t(()=>[c(e("input",{id:"k8s-namespace-new","onUpdate:modelValue":n[6]||(n[6]=o=>s.value.k8sNamespaceSelection=o),type:"text",class:"k-input w-100",placeholder:"your-namespace",required:""},null,512),[[M,s.value.k8sNamespaceSelection]])]),_:1})]),_:1})])):y("",!0),a(),s.value.k8sDataplaneType==="dataplane-type-ingress"?(u(),p("div",Be,[e("p",null,h(r(w))+` natively supports the Kong Ingress. Do you want to deploy
              Kong or another Ingress?
            `,1),a(),i(r(v),{class:"my-6","has-shadow":""},{body:t(()=>[i(k,{"all-inline":"","equal-cols":"","hide-label-col":""},{default:t(()=>[e("label",Pe,[c(e("input",{id:"k8s-ingress-kong","onUpdate:modelValue":n[7]||(n[7]=o=>s.value.k8sIngressBrand=o),class:"k-input",type:"radio",name:"k8s-ingress-brand",value:"kong-ingress",checked:""},null,512),[[_,s.value.k8sIngressBrand]]),a(),Ue]),a(),e("label",Ke,[c(e("input",{id:"k8s-ingress-other","onUpdate:modelValue":n[8]||(n[8]=o=>s.value.k8sIngressBrand=o),class:"k-input",type:"radio",name:"k8s-ingress-brand",value:"other-ingress"},null,512),[[_,s.value.k8sIngressBrand]]),a(),Fe])]),_:1})]),_:1}),a(),i(r(v),{class:"my-6","has-shadow":""},{body:t(()=>[i(k,{title:"Deployments","for-attr":"k8s-deployment-selection"},{default:t(()=>[c(e("input",{id:"k8s-ingress-deployment-new","onUpdate:modelValue":n[9]||(n[9]=o=>s.value.k8sIngressDeployment=o),type:"text",class:"k-input w-100",placeholder:"your-deployment",required:""},null,512),[[M,s.value.k8sIngressDeployment]])]),_:1})]),_:1}),a(),s.value.k8sIngressBrand==="other-ingress"?(u(),x(r(B),{key:0,appearance:"info"},{alertMessage:t(()=>[Ae]),_:1})):y("",!0)])):y("",!0)]),complete:t(()=>[s.value.meshName?(u(),p("div",qe,[I.value===!1?(u(),p("div",ze,[We,a(),je,a(),Oe,a(),i(P,{id:"code-block-kubernetes-command",class:"mt-3",language:"bash",code:r(W)},null,8,["code"])])):y("",!0),a(),i(ne,{"loader-function":L,"should-start":!0,"has-error":b.value,"can-complete":E.value,onHideSiblings:G},{"loading-title":t(()=>[Ge]),"loading-content":t(()=>[Le]),"complete-title":t(()=>[Re]),"complete-content":t(()=>[e("p",null,[a(`
                  Your Dataplane
                  `),s.value.k8sNamespaceSelection?(u(),p("strong",Ye,h(s.value.k8sNamespaceSelection),1)):y("",!0),a(`
                  was found!
                `)]),a(),Qe,a(),e("p",null,[i(r(C),{appearance:"primary",onClick:R},{default:t(()=>[a(`
                    View Your Dataplane
                  `)]),_:1})])]),"error-title":t(()=>[$e]),"error-content":t(()=>[He]),_:1},8,["has-error","can-complete"])])):(u(),x(r(B),{key:1,appearance:"danger"},{alertMessage:t(()=>[Xe]),_:1}))]),dataplane:t(()=>[Je,a(),e("p",null,`
            In `+h(r(w))+`, a Dataplane resource represents a data plane proxy running
            alongside one of your services. Data plane proxies can be added in any Mesh
            that you may have created, and in Kubernetes, they will be auto-injected
            by `+h(r(w))+`.
          `,1)]),example:t(()=>[Ze,a(),ea,a(),i(P,{id:"onboarding-dpp-kubernetes-example",class:"sample-code-block",code:K,language:"yaml"})]),switch:t(()=>[i(le)]),_:1},8,["footer-enabled","next-disabled"])])]))}});const ha=ue(aa,[["__scopeId","data-v-f14bf435"]]);export{ha as default};
