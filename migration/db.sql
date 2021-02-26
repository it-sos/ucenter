drop database if exists ucenter;
create database ucenter charset=utf8mb4 collate=utf8mb4_general_ci;
use ucenter;

drop table if exists user;
create table user (
    id int unsigned not null auto_increment,
    account varchar(32) not null comment '账号',
    password varchar(64) not null comment '密码',
    nickname varchar(16) not null comment '昵称',
    phone varchar(16) not null default '' comment '联系电话',
    expired int unsigned not null default 0 comment '有效期 0=永久',
    is_disabled tinyint unsigned not null default 0 comment '是否禁用状态 1=是；0=否',
    is_deleted tinyint unsigned not null default 0 comment '是否删除状态 1=是；0=否',
    salt varchar(32) not null comment '盐',
    update_time datetime not null on update current_timestamp default current_timestamp comment '更新时间',
    create_time datetime not null default current_timestamp comment '创建时间',
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
    icon varchar(255) not null default '' comment '应用图标',
    is_deleted tinyint unsigned not null default 0 comment '删除状态 1=是；0=否',
    update_time datetime not null on update current_timestamp default current_timestamp comment '更新时间',
    create_time datetime not null default current_timestamp comment '创建时间',
    key idx_appid(appid),
    primary key(id)
) engine=innoDB charset=utf8mb4 comment '应用表';

drop table if exists role;
create table role (
    id int unsigned not null auto_increment,
    app_id int unsigned not null comment '应用表ID',
    name varchar(32) not null comment '角色名',
    info varchar(255) not null comment '描述',
    update_time datetime not null on update current_timestamp default current_timestamp comment '更新时间',
    create_time datetime not null default current_timestamp comment '创建时间',
    key idx_appid(app_id),
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
    update_time datetime not null on update current_timestamp default current_timestamp comment '更新时间',
    create_time datetime not null default current_timestamp comment '创建时间',
    key idx_appid(app_id),
    primary key(id)
) engine=innoDB charset=utf8mb4 comment '权限表';

drop table if exists user_permission;
create table user_permission (
    id int unsigned not null auto_increment,
    user_id int unsigned not null comment '用户表ID',
    permission_id int unsigned not null comment '权限表ID',
    is_forbidden tinyint unsigned not null default 1 comment '是否禁止访问 0=否；1=是',
    update_time datetime not null on update current_timestamp default current_timestamp comment '更新时间',
    create_time datetime not null default current_timestamp comment '创建时间',
    unique uk_userid_permid(user_id, permission_id),
    primary key(id)
) engine=innoDB charset=utf8mb4 comment '用户权限';

drop table if exists role_permission;
create table role_permission (
    id int unsigned not null auto_increment,
    role_id int unsigned not null comment '角色表ID',
    permission_id int unsigned not null comment '权限表ID',
    is_forbidden tinyint unsigned not null default 1 comment '是否禁止访问 0=否；1=是',
    update_time datetime not null on update current_timestamp default current_timestamp comment '更新时间',
    create_time datetime not null default current_timestamp comment '创建时间',
    unique uk_roleid_permid(role_id, permission_id),
    primary key(id)
) engine=innoDB charset=utf8mb4 comment '角色权限';

drop table if exists user_app;
create table user_app (
    id int unsigned not null auto_increment,
    user_id int unsigned not null comment '用户表ID',
    app_id int unsigned not null comment '应用表ID',
    update_time datetime not null on update current_timestamp default current_timestamp comment '更新时间',
    create_time datetime not null default current_timestamp comment '创建时间',
    unique uk_userid_appid(user_id, app_id),
    primary key(id)
) engine=innoDB charset=utf8mb4 comment '用户所属应用表';

drop table if exists user_role;
create table user_role (
    id int unsigned not null auto_increment,
    user_id int unsigned not null comment '用户表ID',
    role_id int unsigned not null comment '角色表ID',
    update_time datetime not null on update current_timestamp default current_timestamp comment '更新时间',
    create_time datetime not null default current_timestamp comment '创建时间',
    unique uk_userid_roleid(user_id, role_id),
    primary key(id)
) engine=innoDB charset=utf8mb4 comment '用户所属角色表';

