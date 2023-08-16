"use strict";
const common_vendor = require("../../common/vendor.js");
const common_assets = require("../../common/assets.js");
const store_user = require("../../store/user.js");
require("../../utils/api/user.js");
require("../../utils/request.js");
require("../../utils/token.js");
const _sfc_main = {
  __name: "login",
  setup(__props) {
    let useStore = store_user.useUserStore();
    common_vendor.ref("");
    let getAccessToken = () => {
      common_vendor.index.login({
        provider: "weixin",
        success(res) {
          let lcode = res.code;
          useStore.userLogin(lcode);
        }
      });
    };
    return (_ctx, _cache) => {
      return {
        a: common_assets._imports_0$2,
        b: common_vendor.o((...args) => common_vendor.unref(getAccessToken) && common_vendor.unref(getAccessToken)(...args))
      };
    };
  }
};
const Component = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__file", "D:/黑马前端基础学习/小程序/test-v3/components/login/login.vue"]]);
wx.createComponent(Component);
