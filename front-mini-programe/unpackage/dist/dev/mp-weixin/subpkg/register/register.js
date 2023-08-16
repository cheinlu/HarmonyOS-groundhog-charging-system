"use strict";
const common_vendor = require("../../common/vendor.js");
const utils_api_user = require("../../utils/api/user.js");
const store_user = require("../../store/user.js");
require("../../utils/request.js");
require("../../utils/token.js");
const _sfc_main = {
  __name: "register",
  setup(__props) {
    let useStore = store_user.useUserStore();
    let phoneNumber = common_vendor.ref("");
    let verificationCode = common_vendor.ref("");
    let codeInputDisabled = common_vendor.ref(true);
    let submitDisabled = common_vendor.ref(true);
    let disableBtn = common_vendor.ref(false);
    let btnText = common_vendor.ref("获取验证码");
    let countdown = common_vendor.ref(5);
    let userInfo = common_vendor.reactive({
      nickName: "",
      avatarUrl: "",
      gender: ""
    });
    let handleInput = () => {
      const reg = /^1[0-9]{10}$/;
      if (reg.test(phoneNumber.value)) {
        disableBtn.value = false;
      } else {
        disableBtn.value = true;
      }
    };
    let sendVerificationCode = () => {
      if (phoneNumber.value == "") {
        common_vendor.index.$showMsg("请输入手机号");
        disableBtn.value = false;
      } else {
        btnText.value = `${countdown.value}s 后重新获取`;
        disableBtn.value = true;
        codeInputDisabled.value = false;
        submitDisabled.value = false;
        const timer = setInterval(() => {
          if (countdown.value <= 1) {
            clearInterval(timer);
            btnText.value = "重新获取";
            countdown.value = 5;
            disableBtn.value = false;
            codeInputDisabled.value = false;
          } else {
            countdown.value--;
            btnText.value = `${countdown.value}s 后重新获取`;
          }
        }, 1e3);
        utils_api_user.requestVerificationCode({
          params: {
            phone: phoneNumber.value
          }
        }).then((res) => {
          if (res.data.code == 0) {
            common_vendor.index.$showMsg("验证码获取成功");
          }
        });
      }
    };
    let handleSubmit = () => {
      if (phoneNumber.value == "" && verificationCode.value !== "666666") {
        common_vendor.index.$showMsg("请输入手机号和验证码");
        return false;
      } else {
        common_vendor.index.getUserProfile({
          desc: "用户信息",
          success(res) {
            useStore.updateUserInfo(res.userInfo);
            userInfo = res.userInfo;
            common_vendor.index.login({
              provider: "weixin",
              success(res2) {
                let regisCode = res2.code;
                let params = {
                  code: regisCode,
                  phone: phoneNumber.value,
                  verifyCode: verificationCode.value,
                  nickname: userInfo.nickName,
                  avatarUrl: userInfo.avatarUrl,
                  gender: userInfo.gender
                };
                utils_api_user.requestRegister(params).then((res3) => {
                  if (res3.data.code == 0) {
                    useStore.updateToken(res3.data.data.token);
                    common_vendor.index.$showMsg("注册成功");
                    common_vendor.index.switchTab({
                      url: "/pages/my/my"
                    });
                  } else {
                    common_vendor.index.$showMsg(res3.data.message);
                  }
                }).catch((err) => {
                  console.error(err);
                  common_vendor.index.$showMsg("服务器出错，请稍后再试");
                });
              }
            });
          }
        });
      }
    };
    return (_ctx, _cache) => {
      return common_vendor.e({
        a: common_vendor.o([($event) => common_vendor.isRef(phoneNumber) ? phoneNumber.value = $event.detail.value : phoneNumber = $event.detail.value, (...args) => common_vendor.unref(handleInput) && common_vendor.unref(handleInput)(...args)]),
        b: common_vendor.unref(phoneNumber),
        c: common_vendor.unref(codeInputDisabled),
        d: common_vendor.unref(verificationCode),
        e: common_vendor.o(($event) => common_vendor.isRef(verificationCode) ? verificationCode.value = $event.detail.value : verificationCode = $event.detail.value),
        f: common_vendor.t(common_vendor.unref(btnText)),
        g: common_vendor.o((...args) => common_vendor.unref(sendVerificationCode) && common_vendor.unref(sendVerificationCode)(...args)),
        h: common_vendor.unref(disableBtn),
        i: common_vendor.unref(submitDisabled)
      }, common_vendor.unref(submitDisabled) ? {} : {
        j: common_vendor.o((...args) => common_vendor.unref(handleSubmit) && common_vendor.unref(handleSubmit)(...args))
      });
    };
  }
};
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__scopeId", "data-v-0d2c45b0"], ["__file", "D:/黑马前端基础学习/小程序/test-v3/subpkg/register/register.vue"]]);
wx.createPage(MiniProgramPage);
