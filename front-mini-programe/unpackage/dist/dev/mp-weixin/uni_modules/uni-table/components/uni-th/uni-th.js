"use strict";
const common_vendor = require("../../../../common/vendor.js");
const _sfc_main = {
  name: "uniTh",
  options: {
    virtualHost: true
  },
  components: {},
  emits: ["sort-change", "filter-change"],
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
    },
    sortable: {
      type: Boolean,
      default: false
    },
    filterType: {
      type: String,
      default: ""
    },
    filterData: {
      type: Array,
      default() {
        return [];
      }
    },
    filterDefaultValue: {
      type: [Array, String],
      default() {
        return "";
      }
    }
  },
  data() {
    return {
      border: false,
      ascending: false,
      descending: false
    };
  },
  computed: {
    // 根据props中的width属性 自动匹配当前th的宽度(px)
    customWidth() {
      if (typeof this.width === "number") {
        return this.width;
      } else if (typeof this.width === "string") {
        let regexHaveUnitPx = new RegExp(/^[1-9][0-9]*px$/g);
        let regexHaveUnitRpx = new RegExp(/^[1-9][0-9]*rpx$/g);
        let regexHaveNotUnit = new RegExp(/^[1-9][0-9]*$/g);
        if (this.width.match(regexHaveUnitPx) !== null) {
          return this.width.replace("px", "");
        } else if (this.width.match(regexHaveUnitRpx) !== null) {
          let numberRpx = Number(this.width.replace("rpx", ""));
          let widthCoe = common_vendor.index.getSystemInfoSync().screenWidth / 750;
          return Math.round(numberRpx * widthCoe);
        } else if (this.width.match(regexHaveNotUnit) !== null) {
          return this.width;
        } else {
          return "";
        }
      } else {
        return "";
      }
    },
    contentAlign() {
      let align = "left";
      switch (this.align) {
        case "left":
          align = "flex-start";
          break;
        case "center":
          align = "center";
          break;
        case "right":
          align = "flex-end";
          break;
      }
      return align;
    }
  },
  created() {
    this.root = this.getTable("uniTable");
    this.rootTr = this.getTable("uniTr");
    this.rootTr.minWidthUpdate(this.customWidth ? this.customWidth : 140);
    this.border = this.root.border;
    this.root.thChildren.push(this);
  },
  methods: {
    sort() {
      if (!this.sortable)
        return;
      this.clearOther();
      if (!this.ascending && !this.descending) {
        this.ascending = true;
        this.$emit("sort-change", { order: "ascending" });
        return;
      }
      if (this.ascending && !this.descending) {
        this.ascending = false;
        this.descending = true;
        this.$emit("sort-change", { order: "descending" });
        return;
      }
      if (!this.ascending && this.descending) {
        this.ascending = false;
        this.descending = false;
        this.$emit("sort-change", { order: null });
      }
    },
    ascendingFn() {
      this.clearOther();
      this.ascending = !this.ascending;
      this.descending = false;
      this.$emit("sort-change", { order: this.ascending ? "ascending" : null });
    },
    descendingFn() {
      this.clearOther();
      this.descending = !this.descending;
      this.ascending = false;
      this.$emit("sort-change", { order: this.descending ? "descending" : null });
    },
    clearOther() {
      this.root.thChildren.map((item) => {
        if (item !== this) {
          item.ascending = false;
          item.descending = false;
        }
        return item;
      });
    },
    ondropdown(e) {
      this.$emit("filter-change", e);
    },
    /**
     * 获取父元素实例
     */
    getTable(name) {
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
function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
  return {
    a: $data.border ? 1 : "",
    b: $options.customWidth + "px",
    c: $props.align
  };
}
const Component = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render], ["__file", "D:/黑马前端基础学习/6.vue/项目/土拨鼠充电/Groundhog-Charging-System/front-mini-programe/uni_modules/uni-table/components/uni-th/uni-th.vue"]]);
wx.createComponent(Component);
