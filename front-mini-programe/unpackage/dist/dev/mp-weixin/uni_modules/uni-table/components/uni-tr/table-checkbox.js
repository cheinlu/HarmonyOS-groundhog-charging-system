"use strict";
const common_vendor = require("../../../../common/vendor.js");
const _sfc_main = {
  name: "TableCheckbox",
  emits: ["checkboxSelected"],
  props: {
    indeterminate: {
      type: Boolean,
      default: false
    },
    checked: {
      type: [Boolean, String],
      default: false
    },
    disabled: {
      type: Boolean,
      default: false
    },
    index: {
      type: Number,
      default: -1
    },
    cellData: {
      type: Object,
      default() {
        return {};
      }
    }
  },
  watch: {
    checked(newVal) {
      if (typeof this.checked === "boolean") {
        this.isChecked = newVal;
      } else {
        this.isChecked = true;
      }
    },
    indeterminate(newVal) {
      this.isIndeterminate = newVal;
    }
  },
  data() {
    return {
      isChecked: false,
      isDisabled: false,
      isIndeterminate: false
    };
  },
  created() {
    if (typeof this.checked === "boolean") {
      this.isChecked = this.checked;
    }
    this.isDisabled = this.disabled;
  },
  methods: {
    selected() {
      if (this.isDisabled)
        return;
      this.isIndeterminate = false;
      this.isChecked = !this.isChecked;
      this.$emit("checkboxSelected", {
        checked: this.isChecked,
        data: this.cellData
      });
    }
  }
};
function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
  return common_vendor.e({
    a: !$props.indeterminate
  }, !$props.indeterminate ? {
    b: $data.isChecked ? 1 : "",
    c: $data.isDisabled ? 1 : ""
  } : {}, {
    d: common_vendor.o((...args) => $options.selected && $options.selected(...args))
  });
}
const Component = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render], ["__file", "C:/Users/123/Desktop/code/lucy-demo/01土拨鼠充电系统/groundhog-charging-system/front-mini-programe/uni_modules/uni-table/components/uni-tr/table-checkbox.vue"]]);
wx.createComponent(Component);
