(global["webpackJsonp"]=global["webpackJsonp"]||[]).push([["subpkg/register/register"],{1692:function(e,t){},"1bac":function(e,t,n){"use strict";n.r(t);var i=n("bf2b"),o=n("9bfc");for(var r in o)["default"].indexOf(r)<0&&function(e){n.d(t,e,(function(){return o[e]}))}(r);n("5584");var u=n("f0c5"),a=Object(u["a"])(o["default"],i["b"],i["c"],!1,null,"034a147c",null,!1,i["a"],void 0);t["default"]=a.exports},2123:function(e,t,n){"use strict";(function(e,t){var i=n("4ea4");n("1692");i(n("66fd"));var o=i(n("1bac"));e.__webpack_require_UNI_MP_PLUGIN__=n,t(o.default)}).call(this,n("bc2e")["default"],n("543d")["createPage"])},5584:function(e,t,n){"use strict";var i=n("e9d4"),o=n.n(i);o.a},"9bfc":function(e,t,n){"use strict";n.r(t);var i=n("d44f"),o=n.n(i);for(var r in i)["default"].indexOf(r)<0&&function(e){n.d(t,e,(function(){return i[e]}))}(r);t["default"]=o.a},bf2b:function(e,t,n){"use strict";n.d(t,"b",(function(){return i})),n.d(t,"c",(function(){return o})),n.d(t,"a",(function(){}));var i=function(){var e=this.$createElement;this._self._c},o=[]},d07a:function(e,t,n){"use strict";(function(e){Object.defineProperty(t,"__esModule",{value:!0}),t.requestVerificationCode=t.requestUserInfo=t.requestRegister=t.requestLogin=void 0;var n=e.$u.http;t.requestLogin=function(e){return n.post("/wx-api/login",e)};t.requestUserInfo=function(){return n.get("/wx-api/me/info")};t.requestVerificationCode=function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return n.get("/wx-api/validcode",e)};t.requestRegister=function(e){return n.post("/wx-api/register",e)}}).call(this,n("543d")["default"])},d44f:function(e,t,n){"use strict";(function(e){Object.defineProperty(t,"__esModule",{value:!0}),t.default=void 0;var i=n("d07a"),o={name:"register",data:function(){return{phoneNumber:"",verificationCode:"",btnText:"获取验证码",disableBtn:!1,submitDisabled:!0,countdown:5,userInfo:{nickname:"",avatarUrl:"",gender:""}}},methods:{handleInput:function(){/^1[0-9]{10}$/.test(this.phoneNumber)?this.disableBtn=!1:this.disableBtn=!0},sendVerificationCode:function(){var t=this;if(""==this.phoneNumber)e.showToast({icon:"none",title:"请输入手机号"}),this.disableBtn=!1;else{this.btnText="".concat(this.countdown,"s 后重新获取"),this.disableBtn=!0;var n=setInterval((function(){t.countdown<=1?(clearInterval(n),t.btnText="重新获取",t.countdown=5,t.disableBtn=!1):(t.countdown--,t.btnText="".concat(t.countdown,"s 后重新获取"))}),1e3);(0,i.requestVerificationCode)({params:{phone:this.phoneNumber}}).then((function(t){0==t.data.code&&e.showToast({icon:"success",title:"验证码"})}))}},handleSubmit:function(){var t=this;if(""==t.phoneNumber&&""==t.verificationCode)return!1;e.getUserProfile({desc:"用户信息",success:function(n){t.userInfo=n.userInfo,e.login({provider:"weixin",success:function(n){var o=n.code,r={code:o,phone:t.phoneNumber,verifyCode:t.verificationCode,nickname:t.userInfo.nickname,avatarUrl:t.userInfo.avatarUrl,gender:t.userInfo.gender};(0,i.requestRegister)(r).then((function(n){console.log("res",n),0==n.data.code&&(e.showToast({icon:"none",title:"注册成功"}),e.setStorageSync("userInfo",t.userInfo),e.switchTab({url:"/pages/my/my"}))}))}})}})}}};t.default=o}).call(this,n("543d")["default"])},e9d4:function(e,t,n){}},[["2123","common/runtime","common/vendor"]]]);