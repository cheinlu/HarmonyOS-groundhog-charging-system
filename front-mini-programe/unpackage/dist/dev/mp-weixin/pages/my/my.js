"use strict";
const common_vendor = require("../../common/vendor.js");
const store_user = require("../../store/user.js");
require("../../utils/api/user.js");
require("../../utils/request.js");
require("../../utils/token.js");
if (!Array) {
  const _easycom_login2 = common_vendor.resolveComponent("login");
  const _easycom_info2 = common_vendor.resolveComponent("info");
  const _easycom_tabbar2 = common_vendor.resolveComponent("tabbar");
  (_easycom_login2 + _easycom_info2 + _easycom_tabbar2)();
}
const _easycom_login = () => "../../components/login/login.js";
const _easycom_info = () => "../../components/info/info.js";
const _easycom_tabbar = () => "../../components/tabbar/tabbar.js";
if (!Math) {
  (_easycom_login + _easycom_info + _easycom_tabbar)();
}
const _sfc_main = {
  __name: "my",
  setup(__props) {
    let useStore = store_user.useUserStore();
    common_vendor.onShow(() => {
      common_vendor.index.hideTabBar({
        animation: false
      });
    });
    return (_ctx, _cache) => {
      return common_vendor.e({
        a: !common_vendor.unref(useStore).token
      }, !common_vendor.unref(useStore).token ? {} : {});
    };
  }
};
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__file", "D:/黑马前端基础学习/6.vue/项目/土拨鼠充电/Groundhog-Charging-System/front-mini-programe/pages/my/my.vue"]]);
wx.createPage(MiniProgramPage);
