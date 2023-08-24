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
    const useStore = store_user.useUserStore();
    common_vendor.ref("");
    const getAccessToken = () => {
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
        b: common_vendor.o(getAccessToken)
      };
    };
  }
};
const Component = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__file", "D:/黑马前端基础学习/6.vue/项目/土拨鼠充电/Groundhog-Charging-System/front-mini-programe/components/login/login.vue"]]);
wx.createComponent(Component);
