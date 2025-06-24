export class TenantValueItem {
  value: string = '';
}

export class TenantsData {
  Data?: TenantItem[]
}

export class TenantItem {
  id?: number;
  name?: string;
  domain?: string;
}

export class LoginData {
  token?: string
  role?: string
}