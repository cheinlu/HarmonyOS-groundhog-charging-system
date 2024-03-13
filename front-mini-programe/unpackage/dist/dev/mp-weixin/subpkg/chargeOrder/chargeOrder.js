"use strict";
const common_vendor = require("../../common/vendor.js");
const common_assets = require("../../common/assets.js");
const utils_api_charge = require("../../utils/api/charge.js");
const utils_time = require("../../utils/time.js");
require("../../utils/request.js");
require("../../env.js");
require("../../store/user.js");
require("../../utils/token.js");
const _sfc_main = {
  __name: "chargeOrder",
  setup(__props) {
    let pageSize = common_vendor.ref(10);
    let pageNo = common_vendor.ref(1);
    let total = common_vendor.ref(0);
    let chargeOrderList = common_vendor.ref([]);
    common_vendor.onMounted(() => {
      getChargeOrderList();
    });
    let getChargeOrderList = async () => {
      let { data: res } = await utils_api_charge.requestChargeOrders(pageNo.value, pageSize.value);
      if (res.code == 0) {
        chargeOrderList.value = chargeOrderList.value.concat(res.data.List);
        pageNo.value = res.data.PageNo;
        pageSize.value = res.data.PageSize;
        total.value = res.data.TotalCount;
      }
    };
    let stopCharge = async (item) => {
      let { data: res } = await utils_api_charge.requestStopCharge(item.id);
      if (res.code == 0) {
        item.state = 1;
        common_vendor.index.$showMsg("已停止充电");
      } else {
        common_vendor.index.$showMsg(res.message);
      }
    };
    let loadMore = () => {
      let totalPage = Math.ceil(total.value / pageSize.value);
      if (pageNo.value >= totalPage) {
        common_vendor.index.$showMsg("已经到最底部了");
        return false;
      } else {
        pageNo.value++;
        getChargeOrderList();
      }
    };
    return (_ctx, _cache) => {
      return common_vendor.e({
        a: common_vendor.unref(chargeOrderList).length === 0
      }, common_vendor.unref(chargeOrderList).length === 0 ? {
        b: common_assets._imports_0$1
      } : {
        c: common_vendor.f(common_vendor.unref(chargeOrderList), (item, k0, i0) => {
          return common_vendor.e({
            a: common_vendor.t(item.orderCode),
            b: item.state === 0
          }, item.state === 0 ? {} : {}, {
            c: common_vendor.t(item.pileCode),
            d: common_vendor.t(item.stationName),
            e: common_vendor.t(item.price),
            f: common_vendor.t(item.state === 0 ? "充电中" : item.state === 1 ? "已完成" : "未付账"),
            g: common_vendor.t(common_vendor.unref(utils_time.formatDate)(item.startAt)),
            h: item.stopAt !== void 0 && item.stopAt !== null && item.state !== 0
          }, item.stopAt !== void 0 && item.stopAt !== null && item.state !== 0 ? {
            i: common_vendor.t(item.state === 0 ? "" : common_vendor.unref(utils_time.formatDate)(item.stopAt))
          } : {}, {
            j: item.state === 0
          }, item.state === 0 ? {
            k: common_vendor.o(($event) => common_vendor.unref(stopCharge)(item), item.id)
          } : {}, {
            l: item.id
          });
        }),
        d: common_vendor.o((...args) => common_vendor.unref(loadMore) && common_vendor.unref(loadMore)(...args))
      });
    };
  }
};
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__scopeId", "data-v-22568b06"], ["__file", "C:/Users/123/Desktop/code/lucy-demo/土拨鼠充电系统/groundhog-charging-system/front-mini-programe/subpkg/chargeOrder/chargeOrder.vue"]]);
wx.createPage(MiniProgramPage);
