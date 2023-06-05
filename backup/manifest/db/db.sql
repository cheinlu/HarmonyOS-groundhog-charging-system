-- DROP TABLE IF EXISTS user;
CREATE TABLE IF NOT EXISTS user(
    id INTEGER PRIMARY KEY AUTOINCREMENT,  -- '自增ID'
    passport  varchar(45) NOT NULL unique, --  'User Passport'
    password  varchar(45) NOT NULL, --  'User Password'
    nickname  varchar(45) NOT NULL, --  'User Nickname'
    create_at datetime(0) DEFAULT NULL, --  'Created Time'
    update_at datetime(0) DEFAULT NULL --  'Updated Time'
);

REPLACE INTO user
(id, passport, password, nickname, create_at, update_at) 
VALUES 
(1, 'admin', '8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918', '管理员', datetime('now'), datetime('now')),
(2, 'guest', '84983c60f7daadc1cb8698621f802c0d9f9a3c3c295c810748fb048115c186ec', '访客', datetime('now'), datetime('now'));

-- DROP TABLE IF EXISTS station;
CREATE TABLE IF NOT EXISTS station(
    id INTEGER PRIMARY KEY AUTOINCREMENT,  -- '自增ID'
    name  varchar(45) NOT NULL unique, --  '充电站名'
    address  varchar(45) NOT NULL, --  '充电站地址'
    coordinate  varchar(45) NOT NULL, --  '坐标，格式 x坐标,y坐标，如113.906737,22.568444'
    create_at datetime(0) DEFAULT NULL, --  'Created Time'
    update_at datetime(0) DEFAULT NULL --  'Updated Time'
);

REPLACE INTO station(id, name, address, coordinate,create_at, update_at) VALUES (1, '宝安新安充电桩', '深圳市宝安区新安节点17区3号', '113.906737,22.568444', datetime('now'), datetime('now'));

-- DROP TABLE IF EXISTS pile;
CREATE TABLE IF NOT EXISTS pile(
    id INTEGER PRIMARY KEY AUTOINCREMENT,  -- '自增ID'
    code  varchar(6) NOT NULL, --  '充电桩号'
    station_id INTEGER NOT NULL, --  '充电站Id'
    state INTEGER NOT NULL, --  '0:空闲中,1:充电中,2:故障'
    create_at datetime(0) DEFAULT NULL, --  'Created Time'
    update_at datetime(0) DEFAULT NULL, --  'Updated Time'
    UNIQUE (code, station_id)
);

REPLACE INTO pile
(id, code, station_id, state,create_at, update_at) 
VALUES 
(1, '01', 1, 0, datetime('now'), datetime('now')),
(2, '02', 1, 1, datetime('now'), datetime('now')),
(3, '03', 1, 2, datetime('now'), datetime('now'));

-- DROP TABLE IF EXISTS charge_order;
CREATE TABLE IF NOT EXISTS charge_order(
    id INTEGER PRIMARY KEY AUTOINCREMENT,  -- '自增ID'
    order_code  varchar(45) NOT NULL UNIQUE, --  '订单号'
    user_id INTEGER NOT NULL, --  '用户Id'
    station_id INTEGER NOT NULL, --  '充电站Id'
    pile_id INTEGER NOT NULL, --  '充电桩 id'
    start_at datetime(0) DEFAULT NULL, --  '充电开始时间'
    stop_at datetime(0) DEFAULT NULL, --  '充电结束时间'
    state INTEGER NOT NULL, --  '0:充电中,1:已完成,2:未付账'
    price NUMERIC DEFAULT 0, --  ''
    create_at datetime(0) DEFAULT NULL, --  'Created Time'
    update_at datetime(0) DEFAULT NULL --  'Updated Time'
);

REPLACE INTO charge_order
(id, order_code,user_id, station_id, pile_id,start_at, stop_at,state,price,create_at,update_at) 
VALUES 
(1, '202305111249-8745',1, 1, 1, datetime('now'), null, 0,0,datetime('now'),datetime('now')),
(2, '202305150133-6666',1, 1, 1, datetime('now'), datetime('now'), 1,1.0,datetime('now'),datetime('now')),
(3, '202305161253-6666',2, 2, 1, datetime('now'), datetime('now'), 2,1.0,datetime('now'),datetime('now'));

CREATE TABLE IF NOT EXISTS charge_price(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    start_hour int NOT NULL, --  '开始时段'
    end_hour int NOT NULL, --  '结束时段'
    price NUMERIC NOT NULL, --  '最小价格'
    create_at datetime(0) DEFAULT NULL, --  '创建时间'
    update_at datetime(0) DEFAULT NULL --  '更新时间'
);
REPLACE INTO charge_price
(id, start_hour,end_hour,price,create_at,update_at) 
VALUES 
(1, 0,9, 0.7, datetime('now'), datetime('now')),
(2, 9,20, 1.5, datetime('now'), datetime('now')),
(3, 20,24, 0.7, datetime('now'), datetime('now'));

-- DROP TABLE IF EXISTS pay_record;
CREATE TABLE IF NOT EXISTS pay_record(
    id INTEGER PRIMARY KEY AUTOINCREMENT,  -- '自增ID'
    pay_code  varchar(45) NOT NULL UNIQUE, --  '订单号'
    user_id INTEGER NOT NULL, --  '用户Id'
    state INTEGER NOT NULL, --  '0:正在充值,1:充值成功,2:取消充值'
    price NUMERIC DEFAULT 0, --  ''
    pay_at datetime(0) DEFAULT NULL, --  '充值时间'
    create_at datetime(0) DEFAULT NULL, --  'Created Time'
    update_at datetime(0) DEFAULT NULL --  'Updated Time'
);

REPLACE INTO pay_record(id, pay_code,user_id, state, price,pay_at,create_at,update_at) VALUES (1, '202305111249-8745',1, 1, 10.0, datetime('now'),datetime('now'),datetime('now'));

-- DROP TABLE IF EXISTS user_wallet;
CREATE TABLE IF NOT EXISTS user_wallet(
    user_id INTEGER PRIMARY KEY,  -- '自增ID'
    balance  NUMERIC NOT NULL, --  '余额'
    create_at datetime(0) DEFAULT NULL, --  'Created Time'
    update_at datetime(0) DEFAULT NULL --  'Updated Time'
);
REPLACE INTO user_wallet
(user_id, balance,create_at,update_at) 
VALUES 
(1, 10.0, datetime('now'),datetime('now')),
(1, 50.0, datetime('now'),datetime('now'));

-- DROP TABLE IF EXISTS wx_user;
CREATE TABLE IF NOT EXISTS wx_user(
    user_id INTEGER PRIMARY KEY,  -- 'userID'
    open_id  varchar(45) NOT NULL, --  '余额'
    phone_no  varchar(45) DEFAULT NULL, --  '手机号'
    avatar_url  varchar(45) DEFAULT NULL, --  '头像地址'
    nickname  varchar(45) DEFAULT NULL, --  '昵称'
    gender  varchar(45) DEFAULT NULL, --  '性别'
    create_at datetime(0) DEFAULT NULL, --  'Created Time'
    update_at datetime(0) DEFAULT NULL --  'Updated Time'
);
