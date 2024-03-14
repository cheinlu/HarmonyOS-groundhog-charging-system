import http from '@ohos.net.http';
import router from '@ohos.router';

class UserModel {
  baseUrl: string = 'http://127.0.0.1:8000';
  tokenKey: string = 'Token';
  httpRequest: any;

  constructor() {
    this.httpRequest = http.createHttp();
  }

  private async sendRequest(url: string, method: string, data: any, headers: any): Promise<any> {
    try {
      const response = await this.httpRequest.request(url, {
        method: method,
        extraData: data,
        header: headers
      });
      return response;
    } catch (error) {
      console.error('请求失败：', error);
      throw new Error('请求失败');
    }
  }

  async getRegister(name: string, domain: string, adminUsername: string, adminPassword: string, validCode: string): Promise<any> {
    return this.sendRequest(`${this.baseUrl}/saas-api/tenant/register`, 'POST', { 'name': name, 'domain': domain, 'adminUsername': adminUsername, 'adminPassword': adminPassword, 'validCode': validCode }, { 'Content-Type': 'application/json' });
  }

  async getToken(tenant: string, account: string, password: string): Promise<any>{
    const headers = {
      'X-Tenant-ID': tenant,
      'Content-Type': 'application/json'
    }
    return this.sendRequest(`${this.baseUrl}/saas-api/user/login`, 'POST',{ 'username': account, 'password': password }, headers)
  }

  async getTenant(): Promise<any> {
    return this.sendRequest(`${this.baseUrl}/saas-api/tenant/selects`, 'GET', null, null);
  }

  async getUserInfo(): Promise<any> {
    const headers = {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${AppStorage.Get('TOKEN')}`
    };

    return this.sendRequest(`${this.baseUrl}/wx-api/me/info`, 'GET', null, headers);
  }
}

const userModel = new UserModel();
export default userModel as UserModel;
