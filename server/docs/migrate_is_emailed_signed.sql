-- 将 x_system_notice.is_emailed 改为有符号 tinyint，以支持 -1（不发送）状态
-- 新语义：-1不发送 0待发送 1发送中(已进入队列) 2发送成功 3发送失败
ALTER TABLE `x_system_notice`
    MODIFY COLUMN `is_emailed` TINYINT NOT NULL DEFAULT 0
    COMMENT '邮件推送状态: -1不发送 0待发送 1发送中(已进入队列) 2发送成功 3发送失败';
