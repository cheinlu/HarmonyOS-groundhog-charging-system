"use strict";
Object.defineProperty(exports, Symbol.toStringTag, { value: "Module" });
const common_vendor = require("./common/vendor.js");
const uni_modules_uviewPlus_index = require("./uni_modules/uview-plus/index.js");
require("./uni_modules/uview-plus/libs/mixin/mixin.js");
require("./uni_modules/uview-plus/libs/mixin/mpMixin.js");
require("./uni_modules/uview-plus/libs/luch-request/core/Request.js");
require("./uni_modules/uview-plus/libs/luch-request/core/dispatchRequest.js");
require("./uni_modules/uview-plus/libs/luch-request/adapters/index.js");
require("./uni_modules/uview-plus/libs/luch-request/helpers/buildURL.js");
require("./uni_modules/uview-plus/libs/luch-request/utils.js");
require("./uni_modules/uview-plus/libs/luch-request/core/buildFullPath.js");
require("./uni_modules/uview-plus/libs/luch-request/helpers/isAbsoluteURL.js");
require("./uni_modules/uview-plus/libs/luch-request/helpers/combineURLs.js");
require("./uni_modules/uview-plus/libs/luch-request/core/settle.js");
require("./uni_modules/uview-plus/libs/luch-request/core/InterceptorManager.js");
require("./uni_modules/uview-plus/libs/luch-request/core/mergeConfig.js");
require("./uni_modules/uview-plus/libs/luch-request/core/defaults.js");
require("./uni_modules/uview-plus/libs/luch-request/utils/clone.js");
require("./uni_modules/uview-plus/libs/util/route.js");
require("./uni_modules/uview-plus/libs/function/colorGradient.js");
require("./uni_modules/uview-plus/libs/function/test.js");
require("./uni_modules/uview-plus/libs/function/debounce.js");
require("./uni_modules/uview-plus/libs/function/throttle.js");
require("./uni_modules/uview-plus/libs/function/index.js");
require("./uni_modules/uview-plus/libs/function/digit.js");
require("./uni_modules/uview-plus/libs/config/config.js");
require("./uni_modules/uview-plus/libs/config/props.js");
require("./uni_modules/uview-plus/libs/config/props/actionSheet.js");
require("./uni_modules/uview-plus/libs/config/props/album.js");
require("./uni_modules/uview-plus/libs/config/props/alert.js");
require("./uni_modules/uview-plus/libs/config/props/avatar.js");
require("./uni_modules/uview-plus/libs/config/props/avatarGroup.js");
require("./uni_modules/uview-plus/libs/config/props/backtop.js");
require("./uni_modules/uview-plus/libs/config/props/badge.js");
require("./uni_modules/uview-plus/libs/config/props/button.js");
require("./uni_modules/uview-plus/libs/config/props/calendar.js");
require("./uni_modules/uview-plus/libs/config/props/carKeyboard.js");
require("./uni_modules/uview-plus/libs/config/props/cell.js");
require("./uni_modules/uview-plus/libs/config/props/cellGroup.js");
require("./uni_modules/uview-plus/libs/config/props/checkbox.js");
require("./uni_modules/uview-plus/libs/config/props/checkboxGroup.js");
require("./uni_modules/uview-plus/libs/config/props/circleProgress.js");
require("./uni_modules/uview-plus/libs/config/props/code.js");
require("./uni_modules/uview-plus/libs/config/props/codeInput.js");
require("./uni_modules/uview-plus/libs/config/props/col.js");
require("./uni_modules/uview-plus/libs/config/props/collapse.js");
require("./uni_modules/uview-plus/libs/config/props/collapseItem.js");
require("./uni_modules/uview-plus/libs/config/props/columnNotice.js");
require("./uni_modules/uview-plus/libs/config/props/countDown.js");
require("./uni_modules/uview-plus/libs/config/props/countTo.js");
require("./uni_modules/uview-plus/libs/config/props/datetimePicker.js");
require("./uni_modules/uview-plus/libs/config/props/divider.js");
require("./uni_modules/uview-plus/libs/config/props/empty.js");
require("./uni_modules/uview-plus/libs/config/props/form.js");
require("./uni_modules/uview-plus/libs/config/props/formItem.js");
require("./uni_modules/uview-plus/libs/config/props/gap.js");
require("./uni_modules/uview-plus/libs/config/props/grid.js");
require("./uni_modules/uview-plus/libs/config/props/gridItem.js");
require("./uni_modules/uview-plus/libs/config/props/icon.js");
require("./uni_modules/uview-plus/libs/config/props/image.js");
require("./uni_modules/uview-plus/libs/config/props/indexAnchor.js");
require("./uni_modules/uview-plus/libs/config/props/indexList.js");
require("./uni_modules/uview-plus/libs/config/props/input.js");
require("./uni_modules/uview-plus/libs/config/props/keyboard.js");
require("./uni_modules/uview-plus/libs/config/props/line.js");
require("./uni_modules/uview-plus/libs/config/props/lineProgress.js");
require("./uni_modules/uview-plus/libs/config/props/link.js");
require("./uni_modules/uview-plus/libs/config/props/list.js");
require("./uni_modules/uview-plus/libs/config/props/listItem.js");
require("./uni_modules/uview-plus/libs/config/props/loadingIcon.js");
require("./uni_modules/uview-plus/libs/config/props/loadingPage.js");
require("./uni_modules/uview-plus/libs/config/props/loadmore.js");
require("./uni_modules/uview-plus/libs/config/props/modal.js");
require("./uni_modules/uview-plus/libs/config/props/navbar.js");
require("./uni_modules/uview-plus/libs/config/color.js");
require("./uni_modules/uview-plus/libs/config/props/noNetwork.js");
require("./uni_modules/uview-plus/libs/config/props/noticeBar.js");
require("./uni_modules/uview-plus/libs/config/props/notify.js");
require("./uni_modules/uview-plus/libs/config/props/numberBox.js");
require("./uni_modules/uview-plus/libs/config/props/numberKeyboard.js");
require("./uni_modules/uview-plus/libs/config/props/overlay.js");
require("./uni_modules/uview-plus/libs/config/props/parse.js");
require("./uni_modules/uview-plus/libs/config/props/picker.js");
require("./uni_modules/uview-plus/libs/config/props/popup.js");
require("./uni_modules/uview-plus/libs/config/props/radio.js");
require("./uni_modules/uview-plus/libs/config/props/radioGroup.js");
require("./uni_modules/uview-plus/libs/config/props/rate.js");
require("./uni_modules/uview-plus/libs/config/props/readMore.js");
require("./uni_modules/uview-plus/libs/config/props/row.js");
require("./uni_modules/uview-plus/libs/config/props/rowNotice.js");
require("./uni_modules/uview-plus/libs/config/props/scrollList.js");
require("./uni_modules/uview-plus/libs/config/props/search.js");
require("./uni_modules/uview-plus/libs/config/props/section.js");
require("./uni_modules/uview-plus/libs/config/props/skeleton.js");
require("./uni_modules/uview-plus/libs/config/props/slider.js");
require("./uni_modules/uview-plus/libs/config/props/statusBar.js");
require("./uni_modules/uview-plus/libs/config/props/steps.js");
require("./uni_modules/uview-plus/libs/config/props/stepsItem.js");
require("./uni_modules/uview-plus/libs/config/props/sticky.js");
require("./uni_modules/uview-plus/libs/config/props/subsection.js");
require("./uni_modules/uview-plus/libs/config/props/swipeAction.js");
require("./uni_modules/uview-plus/libs/config/props/swipeActionItem.js");
require("./uni_modules/uview-plus/libs/config/props/swiper.js");
require("./uni_modules/uview-plus/libs/config/props/swipterIndicator.js");
require("./uni_modules/uview-plus/libs/config/props/switch.js");
require("./uni_modules/uview-plus/libs/config/props/tabbar.js");
require("./uni_modules/uview-plus/libs/config/props/tabbarItem.js");
require("./uni_modules/uview-plus/libs/config/props/tabs.js");
require("./uni_modules/uview-plus/libs/config/props/tag.js");
require("./uni_modules/uview-plus/libs/config/props/text.js");
require("./uni_modules/uview-plus/libs/config/props/textarea.js");
require("./uni_modules/uview-plus/libs/config/props/toast.js");
require("./uni_modules/uview-plus/libs/config/props/toolbar.js");
require("./uni_modules/uview-plus/libs/config/props/tooltip.js");
require("./uni_modules/uview-plus/libs/config/props/transition.js");
require("./uni_modules/uview-plus/libs/config/props/upload.js");
require("./uni_modules/uview-plus/libs/config/zIndex.js");
require("./uni_modules/uview-plus/libs/function/platform.js");
if (!Math) {
  "./pages/charge/charge.js";
  "./pages/my/my.js";
  "./pages/scan/scan.js";
  "./subpkg/charging/charging.js";
  "./subpkg/chargeOrder/chargeOrder.js";
}
const _sfc_main = {
  onLaunch: function() {
    console.log("App Launch");
  },
  onShow: function() {
    console.log("App Show");
  },
  onHide: function() {
    console.log("App Hide");
  }
};
const App = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__file", "D:/黑马前端基础学习/6.vue/项目/土拨鼠充电/Groundhog-Charging-System/front-mini-programe/App.vue"]]);
common_vendor.index.$showMsg = function(title = "数据请求失败！", duration = 1500) {
  common_vendor.index.showToast({
    title,
    duration,
    icon: "none"
  });
};
function createApp() {
  const app = common_vendor.createSSRApp(App);
  app.use(uni_modules_uviewPlus_index.uviewPlus), app.use(common_vendor.createPinia());
  return {
    app,
    Pinia: common_vendor.Pinia
  };
}
createApp().app.mount("#app");
exports.createApp = createApp;
