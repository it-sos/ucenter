drop database if exists ucenter;
create database ucenter charset=utf8mb4 collate=utf8mb4_general_ci;
use ucenter;

drop table if exists user;
create table user (
    id int unsigned not null auto_increment,
    account varchar(32) not null comment '账号',
    password varchar(32) not null comment '密码',
    nickname varchar(16) not null comment '昵称',
    phone varchar(16) not null default '' comment '联系电话',
    exp int unsigned not null default 0 comment '有效期 0=永久',
    status tinyint unsigned not null default 1 comment '状态 1=激活；2=禁用',
    is_del tinyint unsigned not null default 2 comment '删除状态 1=是；2=否',
    salt varchar(32) not null comment '盐',
    utime timestamp on update current_timestamp comment '更新时间',
    ctime timestamp default current_timestamp comment '创建时间',
    key idx_account(account),
    primary key(id)
) engine=innoDB charset=utf8mb4 comment '用户表';

drop table if exists app;
create table app (
    id int unsigned not null auto_increment,
    appid varchar(32) not null comment '应用ID',
    appsecret varchar(32) not null comment '应用秘钥',
    name varchar(32) not null comment '应用名称',
    info varchar(255) not null comment '应用描述',
    link varchar(255) not null default '' comment '应用链接',
    is_del tinyint unsigned not null default 2 comment '删除状态 1=是；2=否',
    utime timestamp on update current_timestamp comment '更新时间',
    ctime timestamp default current_timestamp comment '创建时间',
    key idx_appid(appid),
    primary key(id)
) engine=innoDB charset=utf8mb4 comment '应用表';

drop table if exists role;
create table role (
    id int unsigned not null auto_increment,
    name varchar(32) not null comment '角色名',
    info varchar(255) not null comment '描述',
    utime timestamp on update current_timestamp comment '更新时间',
    ctime timestamp default current_timestamp comment '创建时间',
    primary key(id)
) engine=innoDB charset=utf8mb4 comment '角色表';

drop table if exists permission;
create table permission (
    id int unsigned not null auto_increment,
    app_id int unsigned not null comment '应用表ID',
    name varchar(32) not null comment '功能名称',
    info varchar(255) not null comment '功能描述',
    router varchar(255) not null comment '路由',
    method varchar(8) not null comment '方法 PUT/DELETE/POST/GET',
    parent_id int unsigned not null default 0 comment '父ID',
    utime timestamp on update current_timestamp comment '更新时间',
    ctime timestamp default current_timestamp comment '创建时间',
    key idx_appid(app_id),
    primary key(id)
) engine=innoDB charset=utf8mb4 comment '权限表';

drop table if exists perm_user;
create table perm_user (
    user_id int unsigned not null comment '用户表ID',
    perm_id int unsigned not null comment '权限表ID',
    unique uniq_userid_permid(user_id, perm_id)
) engine=innoDB charset=utf8mb4 comment '用户权限';

drop table if exists perm_role;
create table perm_role (
    role_id int unsigned not null comment '角色表ID',
    perm_id int unsigned not null comment '权限表ID',
    unique uniq_roleid_permid(role_id, perm_id)
) engine=innoDB charset=utf8mb4 comment '角色权限';

drop table if exists role_user;
create table role_user (
    role_id int unsigned not null comment '角色表ID',
    user_id int unsigned not null comment '用户表ID',
    unique uniq_roleid_userid(role_id, user_id)
) engine=innoDB charset=utf8mb4 comment '角色成员表';