
// 与EntryAbility交互，获取ability当前的数据
export class GlobalContext {
  static readonly CONTEXT = 'context';
  static readonly WANT = 'want';
  //状态栏高度
  static readonly STATUS_BAR_HEIGHT = 'statusBarHeight';
  //导航栏高度
  static readonly NAVIGATION_HEIGHT = 'navigationHeight';

  private constructor() { }
  private static instance: GlobalContext;
  private _objects = new Map<string, Object>();

  public static getContext(): GlobalContext {
    if (!GlobalContext.instance) {
      GlobalContext.instance = new GlobalContext();
    }
    return GlobalContext.instance;
  }

  getObject(value: string): Object | undefined {
    return this._objects.get(value);
  }

  setObject(key: string, objectClass: Object): void {
    this._objects.set(key, objectClass);
  }
}