"use strict";
const common_vendor = require("../../../../common/vendor.js");
const _sfc_main = {
  name: "uniTable",
  options: {
    virtualHost: true
  },
  emits: ["selection-change"],
  props: {
    data: {
      type: Array,
      default() {
        return [];
      }
    },
    // 是否有竖线
    border: {
      type: Boolean,
      default: false
    },
    // 是否显示斑马线
    stripe: {
      type: Boolean,
      default: false
    },
    // 多选
    type: {
      type: String,
      default: ""
    },
    // 没有更多数据
    emptyText: {
      type: String,
      default: "没有更多数据"
    },
    loading: {
      type: Boolean,
      default: false
    },
    rowKey: {
      type: String,
      default: ""
    }
  },
  data() {
    return {
      noData: true,
      minWidth: 0,
      multiTableHeads: []
    };
  },
  watch: {
    loading(val) {
    },
    data(newVal) {
      this.theadChildren;
      if (this.theadChildren) {
        this.theadChildren.rowspan;
      }
      this.noData = false;
    }
  },
  created() {
    this.trChildren = [];
    this.thChildren = [];
    this.theadChildren = null;
    this.backData = [];
    this.backIndexData = [];
  },
  methods: {
    isNodata() {
      this.theadChildren;
      let rowspan = 1;
      if (this.theadChildren) {
        rowspan = this.theadChildren.rowspan;
      }
      this.noData = this.trChildren.length - rowspan <= 0;
    },
    /**
     * 选中所有
     */
    selectionAll() {
      let startIndex = 1;
      let theadChildren = this.theadChildren;
      if (!this.theadChildren) {
        theadChildren = this.trChildren[0];
      } else {
        startIndex = theadChildren.rowspan - 1;
      }
      let isHaveData = this.data && this.data.length > 0;
      theadChildren.checked = true;
      theadChildren.indeterminate = false;
      this.trChildren.forEach((item, index) => {
        if (!item.disabled) {
          item.checked = true;
          if (isHaveData && item.keyValue) {
            const row = this.data.find((v) => v[this.rowKey] === item.keyValue);
            if (!this.backData.find((v) => v[this.rowKey] === row[this.rowKey])) {
              this.backData.push(row);
            }
          }
          if (index > startIndex - 1 && this.backIndexData.indexOf(index - startIndex) === -1) {
            this.backIndexData.push(index - startIndex);
          }
        }
      });
      this.$emit("selection-change", {
        detail: {
          value: this.backData,
          index: this.backIndexData
        }
      });
    },
    /**
     * 用于多选表格，切换某一行的选中状态，如果使用了第二个参数，则是设置这一行选中与否（selected 为 true 则选中）
     */
    toggleRowSelection(row, selected) {
      row = [].concat(row);
      this.trChildren.forEach((item, index) => {
        const select = row.findIndex((v) => {
          if (typeof v === "number") {
            return v === index - 1;
          } else {
            return v[this.rowKey] === item.keyValue;
          }
        });
        let ischeck = item.checked;
        if (select !== -1) {
          if (typeof selected === "boolean") {
            item.checked = selected;
          } else {
            item.checked = !item.checked;
          }
          if (ischeck !== item.checked) {
            this.check(item.rowData || item, item.checked, item.rowData ? item.keyValue : null, true);
          }
        }
      });
      this.$emit("selection-change", {
        detail: {
          value: this.backData,
          index: this.backIndexData
        }
      });
    },
    /**
     * 用于多选表格，清空用户的选择
     */
    clearSelection() {
      let theadChildren = this.theadChildren;
      if (!this.theadChildren) {
        theadChildren = this.trChildren[0];
      }
      theadChildren.checked = false;
      theadChildren.indeterminate = false;
      this.trChildren.forEach((item) => {
        item.checked = false;
      });
      this.backData = [];
      this.backIndexData = [];
      this.$emit("selection-change", {
        detail: {
          value: [],
          index: []
        }
      });
    },
    /**
     * 用于多选表格，切换所有行的选中状态
     */
    toggleAllSelection() {
      let list = [];
      let startIndex = 1;
      let theadChildren = this.theadChildren;
      if (!this.theadChildren) {
        theadChildren = this.trChildren[0];
      } else {
        startIndex = theadChildren.rowspan - 1;
      }
      this.trChildren.forEach((item, index) => {
        if (!item.disabled) {
          if (index > startIndex - 1) {
            list.push(index - startIndex);
          }
        }
      });
      this.toggleRowSelection(list);
    },
    /**
     * 选中\取消选中
     * @param {Object} child
     * @param {Object} check
     * @param {Object} rowValue
     */
    check(child, check, keyValue, emit) {
      let theadChildren = this.theadChildren;
      if (!this.theadChildren) {
        theadChildren = this.trChildren[0];
      }
      let childDomIndex = this.trChildren.findIndex((item, index) => child === item);
      if (childDomIndex < 0) {
        childDomIndex = this.data.findIndex((v) => v[this.rowKey] === keyValue) + 1;
      }
      this.trChildren.filter((v) => !v.disabled && v.keyValue).length;
      if (childDomIndex === 0) {
        check ? this.selectionAll() : this.clearSelection();
        return;
      }
      if (check) {
        if (keyValue) {
          this.backData.push(child);
        }
        this.backIndexData.push(childDomIndex - 1);
      } else {
        const index = this.backData.findIndex((v) => v[this.rowKey] === keyValue);
        const idx = this.backIndexData.findIndex((item) => item === childDomIndex - 1);
        if (keyValue) {
          this.backData.splice(index, 1);
        }
        this.backIndexData.splice(idx, 1);
      }
      const domCheckAll = this.trChildren.find((item, index) => index > 0 && !item.checked && !item.disabled);
      if (!domCheckAll) {
        theadChildren.indeterminate = false;
        theadChildren.checked = true;
      } else {
        theadChildren.indeterminate = true;
        theadChildren.checked = false;
      }
      if (this.backIndexData.length === 0) {
        theadChildren.indeterminate = false;
      }
      if (!emit) {
        this.$emit("selection-change", {
          detail: {
            value: this.backData,
            index: this.backIndexData
          }
        });
      }
    }
  }
};
function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
  return common_vendor.e({
    a: $data.noData
  }, $data.noData ? {
    b: common_vendor.t($props.emptyText),
    c: $props.border ? 1 : ""
  } : {}, {
    d: $props.loading
  }, $props.loading ? {
    e: $props.border ? 1 : ""
  } : {}, {
    f: $data.minWidth + "px",
    g: $props.stripe ? 1 : "",
    h: $props.border ? 1 : "",
    i: !$data.noData ? 1 : ""
  });
}
const Component = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render], ["__file", "C:/Users/123/Desktop/code/lucy-demo/01土拨鼠充电系统/groundhog-charging-system/front-mini-programe/uni_modules/uni-table/components/uni-table/uni-table.vue"]]);
wx.createComponent(Component);
