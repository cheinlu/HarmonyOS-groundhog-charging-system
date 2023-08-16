"use strict";
const common_vendor = require("../../../../common/vendor.js");
const _sfc_main = {
  name: "uniTd",
  options: {
    virtualHost: true
  },
  props: {
    width: {
      type: [String, Number],
      default: ""
    },
    align: {
      type: String,
      default: "left"
    },
    rowspan: {
      type: [Number, String],
      default: 1
    },
    colspan: {
      type: [Number, String],
      default: 1
    }
  },
  data() {
    return {
      border: false
    };
  },
  created() {
    this.root = this.getTable();
    this.border = this.root.border;
  },
  methods: {
    /**
     * 获取父元素实例
     */
    getTable() {
      let parent = this.$parent;
      let parentName = parent.$options.name;
      while (parentName !== "uniTable") {
        parent = parent.$parent;
        if (!parent)
          return false;
        parentName = parent.$options.name;
      }
      return parent;
    }
  }
};
function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
  return {
    a: $data.border ? 1 : "",
    b: $props.width + "px",
    c: $props.align
  };
}
const Component = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render], ["__file", "D:/黑马前端基础学习/小程序/test-v3/uni_modules/uni-table/components/uni-td/uni-td.vue"]]);
wx.createComponent(Component);
