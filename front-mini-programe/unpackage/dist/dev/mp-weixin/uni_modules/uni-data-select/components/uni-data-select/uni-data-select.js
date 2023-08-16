"use strict";
const common_vendor = require("../../../../common/vendor.js");
const _sfc_main = {
  name: "uni-data-select",
  mixins: [common_vendor.Ds.mixinDatacom || {}],
  props: {
    localdata: {
      type: Array,
      default() {
        return [];
      }
    },
    value: {
      type: [String, Number],
      default: ""
    },
    modelValue: {
      type: [String, Number],
      default: ""
    },
    label: {
      type: String,
      default: ""
    },
    placeholder: {
      type: String,
      default: "请选择"
    },
    emptyTips: {
      type: String,
      default: "无选项"
    },
    clear: {
      type: Boolean,
      default: true
    },
    defItem: {
      type: Number,
      default: 0
    },
    disabled: {
      type: Boolean,
      default: false
    },
    // 格式化输出 用法 field="_id as value, version as text, uni_platform as label" format="{label} - {text}"
    format: {
      type: String,
      default: ""
    }
  },
  data() {
    return {
      showSelector: false,
      current: "",
      mixinDatacomResData: [],
      apps: [],
      channels: [],
      cacheKey: "uni-data-select-lastSelectedValue"
    };
  },
  created() {
    this.debounceGet = this.debounce(() => {
      this.query();
    }, 300);
    if (this.collection && !this.localdata.length) {
      this.debounceGet();
    }
  },
  computed: {
    typePlaceholder() {
      const text = {
        "opendb-stat-app-versions": "版本",
        "opendb-app-channels": "渠道",
        "opendb-app-list": "应用"
      };
      const common = this.placeholder;
      const placeholder = text[this.collection];
      return placeholder ? common + placeholder : common;
    },
    valueCom() {
      return this.modelValue;
    }
  },
  watch: {
    localdata: {
      immediate: true,
      handler(val, old) {
        if (Array.isArray(val) && old !== val) {
          this.mixinDatacomResData = val;
        }
      }
    },
    valueCom(val, old) {
      this.initDefVal();
    },
    mixinDatacomResData: {
      immediate: true,
      handler(val) {
        if (val.length) {
          this.initDefVal();
        }
      }
    }
  },
  methods: {
    debounce(fn, time = 100) {
      let timer = null;
      return function(...args) {
        if (timer)
          clearTimeout(timer);
        timer = setTimeout(() => {
          fn.apply(this, args);
        }, time);
      };
    },
    // 执行数据库查询
    query() {
      this.mixinDatacomEasyGet();
    },
    // 监听查询条件变更事件
    onMixinDatacomPropsChange() {
      if (this.collection) {
        this.debounceGet();
      }
    },
    initDefVal() {
      let defValue = "";
      if ((this.valueCom || this.valueCom === 0) && !this.isDisabled(this.valueCom)) {
        defValue = this.valueCom;
      } else {
        let strogeValue;
        if (this.collection) {
          strogeValue = this.getCache();
        }
        if (strogeValue || strogeValue === 0) {
          defValue = strogeValue;
        } else {
          let defItem = "";
          if (this.defItem > 0 && this.defItem <= this.mixinDatacomResData.length) {
            defItem = this.mixinDatacomResData[this.defItem - 1].value;
          }
          defValue = defItem;
        }
        if (defValue || defValue === 0) {
          this.emit(defValue);
        }
      }
      const def = this.mixinDatacomResData.find((item) => item.value === defValue);
      this.current = def ? this.formatItemName(def) : "";
    },
    /**
     * @param {[String, Number]} value
     * 判断用户给的 value 是否同时为禁用状态
     */
    isDisabled(value) {
      let isDisabled = false;
      this.mixinDatacomResData.forEach((item) => {
        if (item.value === value) {
          isDisabled = item.disable;
        }
      });
      return isDisabled;
    },
    clearVal() {
      this.emit("");
      if (this.collection) {
        this.removeCache();
      }
    },
    change(item) {
      if (!item.disable) {
        this.showSelector = false;
        this.current = this.formatItemName(item);
        this.emit(item.value);
      }
    },
    emit(val) {
      this.$emit("input", val);
      this.$emit("update:modelValue", val);
      this.$emit("change", val);
      if (this.collection) {
        this.setCache(val);
      }
    },
    toggleSelector() {
      if (this.disabled) {
        return;
      }
      this.showSelector = !this.showSelector;
    },
    formatItemName(item) {
      let {
        text,
        value,
        channel_code
      } = item;
      channel_code = channel_code ? `(${channel_code})` : "";
      if (this.format) {
        let str = "";
        str = this.format;
        for (let key in item) {
          str = str.replace(new RegExp(`{${key}}`, "g"), item[key]);
        }
        return str;
      } else {
        return this.collection.indexOf("app-list") > 0 ? `${text}(${value})` : text ? text : `未命名${channel_code}`;
      }
    },
    // 获取当前加载的数据
    getLoadData() {
      return this.mixinDatacomResData;
    },
    // 获取当前缓存key
    getCurrentCacheKey() {
      return this.collection;
    },
    // 获取缓存
    getCache(name = this.getCurrentCacheKey()) {
      let cacheData = common_vendor.index.getStorageSync(this.cacheKey) || {};
      return cacheData[name];
    },
    // 设置缓存
    setCache(value, name = this.getCurrentCacheKey()) {
      let cacheData = common_vendor.index.getStorageSync(this.cacheKey) || {};
      cacheData[name] = value;
      common_vendor.index.setStorageSync(this.cacheKey, cacheData);
    },
    // 删除缓存
    removeCache(name = this.getCurrentCacheKey()) {
      let cacheData = common_vendor.index.getStorageSync(this.cacheKey) || {};
      delete cacheData[name];
      common_vendor.index.setStorageSync(this.cacheKey, cacheData);
    }
  }
};
if (!Array) {
  const _easycom_uni_icons2 = common_vendor.resolveComponent("uni-icons");
  _easycom_uni_icons2();
}
const _easycom_uni_icons = () => "../../../uni-icons/components/uni-icons/uni-icons.js";
if (!Math) {
  _easycom_uni_icons();
}
function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
  return common_vendor.e({
    a: $props.label
  }, $props.label ? {
    b: common_vendor.t($props.label + "：")
  } : {}, {
    c: $data.current
  }, $data.current ? {
    d: common_vendor.t($data.current)
  } : {
    e: common_vendor.t($options.typePlaceholder)
  }, {
    f: $data.current && $props.clear && !$props.disabled
  }, $data.current && $props.clear && !$props.disabled ? {
    g: common_vendor.p({
      type: "clear",
      color: "#c0c4cc",
      size: "24"
    }),
    h: common_vendor.o((...args) => $options.clearVal && $options.clearVal(...args))
  } : {
    i: common_vendor.p({
      type: $data.showSelector ? "top" : "bottom",
      size: "14",
      color: "#999"
    })
  }, {
    j: common_vendor.o((...args) => $options.toggleSelector && $options.toggleSelector(...args)),
    k: $data.showSelector
  }, $data.showSelector ? {
    l: common_vendor.o((...args) => $options.toggleSelector && $options.toggleSelector(...args))
  } : {}, {
    m: $data.showSelector
  }, $data.showSelector ? common_vendor.e({
    n: $data.mixinDatacomResData.length === 0
  }, $data.mixinDatacomResData.length === 0 ? {
    o: common_vendor.t($props.emptyTips)
  } : {
    p: common_vendor.f($data.mixinDatacomResData, (item, index, i0) => {
      return {
        a: common_vendor.t($options.formatItemName(item)),
        b: item.disable ? 1 : "",
        c: index,
        d: common_vendor.o(($event) => $options.change(item), index)
      };
    })
  }) : {}, {
    q: $props.disabled ? 1 : "",
    r: $data.current ? 1 : ""
  });
}
const Component = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render], ["__file", "D:/黑马前端基础学习/小程序/test-v3/uni_modules/uni-data-select/components/uni-data-select/uni-data-select.vue"]]);
wx.createComponent(Component);
