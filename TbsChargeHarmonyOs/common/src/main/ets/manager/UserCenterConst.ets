import { StorageConst } from "./StorageConst"
import { UserInfoData } from "../model/UserInfoModel"
import { commonStorage } from "../utils"

class UserCenterConst {
  userInfo: UserInfoData | null = null

  getUserInfoData(): UserInfoData | null {
    if (this.userInfo) {
      return this.userInfo;
    }

    const value = commonStorage.getStringValueSync(StorageConst.USER_INFO, '');

    // 如果已经是对象，直接使用
    if (value && typeof value === 'object') {
      this.userInfo = value as UserInfoData;
      return this.userInfo;
    }

    // 如果是字符串，尝试解析
    if (typeof value === 'string' && value) {
      const p = JsonParse<UserInfoData>(value);
      if (p) {
        this.userInfo = p;
        return this.userInfo;
      }
    }

    return null;
  }
}

function JsonParse<T>(text: string): T | undefined {
  try {
    return JSON.parse(text);
  } catch (e) {
    console.error(`JSON解析错误: ${text}, 错误信息: ${e}`);
    return undefined;
  }
}

export default new UserCenterConst()