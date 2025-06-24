import dataPreferences from '@ohos.data.preferences';
import preferences from '@ohos.data.preferences';
import { HashMap } from '@kit.ArkTS';
import { common } from '@kit.AbilityKit';
import { GlobalContext } from './GlobalContext';
import { BusinessError } from '@kit.BasicServicesKit';

class Storage{
  private name: string = 'app_storage'
  private preferences: dataPreferences.Preferences | null = null
  private listenerKeyMap = new HashMap<string, HashMap<string, Callback<string>>>()

  constructor(name: string) {
    this.name = name;
  }

  initStorage() {
    const context = GlobalContext.getContext().getObject('context') as common.UIAbilityContext
    const options: dataPreferences.Options = { name: this.name }
    this.preferences = dataPreferences.getPreferencesSync(context, options)
    this.preferences?.on('change', (key: string) => {
      const pageMap = this.listenerKeyMap.get(key)
      pageMap?.forEach((callBack: Callback<string>) => {
        callBack(key)
      })
    })
  }

  /**
   * 存储key,value到preference缓存里面
   * @param key
   * @param value
   */
  saveValueSync(key: string, value: preferences.ValueType) {
    this.preferences?.putSync(key, value)
    this.preferences?.flush((err: BusinessError) => {
      if (err) {
        console.error(`${this.name} file flush error :${err.message ?? ""}`)
      }
    })
  }

  /**
   * 根据key删除值
   * @param key
   */
  deleteValueSync(key: string) {
    this.preferences?.deleteSync(key)
    this.preferences?.flush((err: BusinessError) => {
      if (err) {
        console.error(`${this.name} file flush error :${err.message ?? ""}`)
      }
    })
  }

  /**
   * 适用于普通函数获取
   * @param key
   * @param defaultValue
   * @returns string
   */
  getStringValueSync(key: string, defaultValue: string) {
    return this.preferences?.getSync(key, defaultValue) as string
  }

  /**
   * 适用于普通函数获取
   * @param key
   * @param defaultValue
   * @returns Array<string>
   */
  getArrayStringValueSync(key: string, defaultValue: Array<string>) {
    return this.preferences?.getSync(key, defaultValue) as Array<string>
  }
}
export const commonStorage = new Storage('common_storage')

export function initDataPreference() {
  commonStorage.initStorage()
}