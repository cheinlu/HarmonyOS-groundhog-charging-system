"use strict";
const common_vendor = require("../../../../common/vendor.js");
const tableCheckbox = () => "./table-checkbox.js";
const _sfc_main = {
  name: "uniTr",
  components: { tableCheckbox },
  props: {
    disabled: {
      type: Boolean,
      default: false
    },
    keyValue: {
      type: [String, Number],
      default: ""
    }
  },
  options: {
    virtualHost: true
  },
  data() {
    return {
      value: false,
      border: false,
      selection: false,
      widthThArr: [],
      ishead: true,
      checked: false,
      indeterminate: false
    };
  },
  created() {
    this.root = this.getTable();
    this.head = this.getTable("uniThead");
    if (this.head) {
      this.ishead = false;
      this.head.init(this);
    }
    this.border = this.root.border;
    this.selection = this.root.type;
    this.root.trChildren.push(this);
    const rowData = this.root.data.find((v) => v[this.root.rowKey] === this.keyValue);
    if (rowData) {
      this.rowData = rowData;
    }
    this.root.isNodata();
  },
  mounted() {
    if (this.widthThArr.length > 0) {
      const selectionWidth = this.selection === "selection" ? 50 : 0;
      this.root.minWidth = this.widthThArr.reduce((a, b) => Number(a) + Number(b)) + selectionWidth;
    }
  },
  unmounted() {
    const index = this.root.trChildren.findIndex((i) => i === this);
    this.root.trChildren.splice(index, 1);
    this.root.isNodata();
  },
  methods: {
    minWidthUpdate(width) {
      this.widthThArr.push(width);
    },
    // 选中
    checkboxSelected(e) {
      let rootData = this.root.data.find((v) => v[this.root.rowKey] === this.keyValue);
      this.checked = e.checked;
      this.root.check(rootData || this, e.checked, rootData ? this.keyValue : null);
    },
    change(e) {
      this.root.trChildren.forEach((item) => {
        if (item === this) {
          this.root.check(this, e.detail.value.length > 0 ? true : false);
        }
      });
    },
    /**
     * 获取父元素实例
     */
    getTable(name = "uniTable") {
      let parent = this.$parent;
      let parentName = parent.$options.name;
      while (parentName !== name) {
        parent = parent.$parent;
        if (!parent)
          return false;
        parentName = parent.$options.name;
      }
      return parent;
    }
  }
};
if (!Array) {
  const _component_table_checkbox = common_vendor.resolveComponent("table-checkbox");
  _component_table_checkbox();
}
function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
  return common_vendor.e({
    a: $data.selection === "selection"
  }, $data.selection === "selection" ? {
    b: common_vendor.o($options.checkboxSelected),
    c: common_vendor.p({
      checked: $data.checked,
      indeterminate: $data.indeterminate,
      disabled: $props.disabled
    }),
    d: $data.border ? 1 : ""
  } : {});
}
const Component = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render], ["__file", "D:/黑马前端基础学习/6.vue/项目/土拨鼠充电/Groundhog-Charging-System/front-mini-programe/uni_modules/uni-table/components/uni-tr/uni-tr.vue"]]);
wx.createComponent(Component);
