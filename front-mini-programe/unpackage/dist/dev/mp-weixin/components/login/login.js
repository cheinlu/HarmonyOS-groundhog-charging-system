"use strict";
const common_vendor = require("../../common/vendor.js");
const common_assets = require("../../common/assets.js");
const utils_api_user = require("../../utils/api/user.js");
const utils_token = require("../../utils/token.js");
const store_user = require("../../store/user.js");
require("../../utils/request.js");
require("../../env.js");
const _sfc_main = {
  __name: "login",
  setup(__props) {
    let useStore = store_user.useUserStore();
    const getLogin = () => {
      common_vendor.index.login({
        provider: "weixin",
        success(res) {
          const code = { code: res.code };
          utils_api_user.requestLogin(code).then((lres) => {
            if (lres.data.code == 0) {
              useStore.token = lres.data.data.token;
              utils_token.SET_TOKEN(lres.data.data.token);
            } else {
              common_vendor.index.showModal({
                title: "温馨提示",
                content: "微信授权登录后才能完整使用小程序功能",
                confirmText: "同意",
                cancelText: "拒绝",
                success: function(res2) {
                  if (res2.confirm) {
                    getRegister();
                  } else if (res2.cancel) {
                    console.log("点击了拒绝按钮，不做操作");
                  }
                }
              });
            }
          });
        }
      });
    };
    const getRegister = () => {
      common_vendor.index.getUserProfile({
        desc: "用于完善会员信息",
        lang: "zh_CN",
        success(res) {
          useStore.updateUserInfo(res.userInfo);
          utils_token.SET_INFO(res.userInfo);
          let userInfo = res.userInfo;
          common_vendor.index.login({
            provider: "weixin",
            success(res2) {
              let regisCode = res2.code;
              let params = {
                code: regisCode,
                nickname: userInfo.nickName,
                avatarUrl: userInfo.avatarUrl,
                gender: userInfo.gender
              };
              utils_api_user.requestRegister(params).then((res3) => {
                if (res3.data.code == 0) {
                  useStore.updateToken(res3.data.data.token);
                  getLogin();
                }
              }).catch((err) => {
                console.error(err);
                common_vendor.index.$showMsg("服务器出错，请稍后再试");
              });
            }
          });
        }
      });
    };
    return (_ctx, _cache) => {
      return {
        a: common_assets._imports_0$2,
        b: common_vendor.o(getLogin)
      };
    };
  }
};
const Component = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__file", "C:/Users/123/Desktop/code/lucy-demo/01土拨鼠充电系统/groundhog-charging-system-github/HarmonyOS-groundhog-charging-system/front-mini-programe/components/login/login.vue"]]);
wx.createComponent(Component);
