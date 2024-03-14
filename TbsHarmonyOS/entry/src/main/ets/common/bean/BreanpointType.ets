//定义标记类型
declare interface BreakpointTypeOptions<T>{
  sm?:T,
  md?:T,
  lg?:T
}

export default class BreakpointType<T>{
  options: BreakpointTypeOptions<T>

  constructor(options: BreakpointTypeOptions<T>) {
    this.options = options
  }

  getValue(breakpoint: string): T{
    return this.options[breakpoint]
  }
}