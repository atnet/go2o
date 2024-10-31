
TRUNCATE TABLE "ad_list";
TRUNCATE TABLE "ad_userset";
TRUNCATE TABLE "article_category";
TRUNCATE TABLE "article_list";
TRUNCATE TABLE "comm_qr_template";
TRUNCATE TABLE "dlv_merchant_bind";
TRUNCATE TABLE "arc_page";
TRUNCATE TABLE "gs_member_price";
TRUNCATE TABLE "item_sku";
TRUNCATE TABLE "item_info";
TRUNCATE TABLE "item_snapshot";
TRUNCATE TABLE "item_trade_snapshot";
TRUNCATE TABLE "mch_account";
TRUNCATE TABLE "mch_api_info";
TRUNCATE TABLE "mch_balance_log";
TRUNCATE TABLE "mch_buyer_group";
TRUNCATE TABLE "mch_day_chart";
TRUNCATE TABLE "mch_enterprise_info";
TRUNCATE TABLE "mch_merchant";
TRUNCATE TABLE "mch_online_shop";
TRUNCATE TABLE "mch_offline_shop";


TRUNCATE TABLE "mm_member";
TRUNCATE TABLE "mm_profile";
TRUNCATE TABLE "mm_integral_log";
TRUNCATE TABLE "mm_balance_log";
TRUNCATE TABLE "mm_wallet_log";
TRUNCATE TABLE "mm_flow_log";
TRUNCATE TABLE "mm_account";
TRUNCATE TABLE "mm_deliver_addr";
TRUNCATE TABLE "mm_relation";
TRUNCATE TABLE "mm_cert_info";
TRUNCATE TABLE "mm_favorite";
TRUNCATE TABLE "mm_bank_card";
TRUNCATE TABLE "mm_receipts_code";
TRUNCATE TABLE "mm_levelup";
TRUNCATE TABLE "mm_buyer_group";
TRUNCATE TABLE "mm_lock_info";
TRUNCATE TABLE "mm_lock_history";
TRUNCATE TABLE "mm_oauth_account";
TRUNCATE TABLE "mm_block_list";



TRUNCATE TABLE "wal_wallet";
TRUNCATE TABLE "wal_wallet_log";

TRUNCATE TABLE "sale_sub_order";
TRUNCATE TABLE "sale_order_item";
TRUNCATE TABLE "pt_order_coupon";
TRUNCATE TABLE "pt_order_pb";
TRUNCATE TABLE "sale_order_log";
TRUNCATE TABLE "sale_cart";
TRUNCATE TABLE "sale_cart_item";
TRUNCATE TABLE "order_list";
TRUNCATE TABLE "order_wholesale_order";
TRUNCATE TABLE "order_wholesale_item";
TRUNCATE TABLE "order_trade_order";

TRUNCATE TABLE "sale_after_order";
TRUNCATE TABLE "sale_return";
TRUNCATE TABLE "sale_exchange";
TRUNCATE TABLE "sale_refund";

TRUNCATE TABLE "express_provider";
TRUNCATE TABLE "mch_express_template";
TRUNCATE TABLE "express_area_set";

TRUNCATE TABLE "ship_order";
TRUNCATE TABLE "ship_item";

TRUNCATE TABLE "pro_product";
TRUNCATE TABLE "item_info";
TRUNCATE TABLE "item_sku";
TRUNCATE TABLE "prod_category";
TRUNCATE TABLE "pro_model";
TRUNCATE TABLE "pro_model_brand";
TRUNCATE TABLE "pro_brand";
TRUNCATE TABLE "pro_attr";
TRUNCATE TABLE "pro_attr_item";
TRUNCATE TABLE "pro_spec";
TRUNCATE TABLE "pro_spec_item";
TRUNCATE TABLE "pro_attr_info";
TRUNCATE TABLE "item_snapshot";
TRUNCATE TABLE "item_trade_snapshot";
TRUNCATE TABLE "gs_sale_label";
TRUNCATE TABLE "gs_member_price";


TRUNCATE TABLE "mch_merchant";
TRUNCATE TABLE "mch_enterprise_info";
TRUNCATE TABLE "mch_api_info";
TRUNCATE TABLE "mch_shop";
TRUNCATE TABLE "mch_online_shop";
TRUNCATE TABLE "mch_offline_shop";
--TRUNCATE TABLE "mch_sale_conf";
TRUNCATE TABLE "mch_trade_conf";
TRUNCATE TABLE "pt_member_level";
TRUNCATE TABLE "mch_account";
TRUNCATE TABLE "mch_balance_log";
TRUNCATE TABLE "mch_day_chart";
TRUNCATE TABLE "mch_sign_up";
TRUNCATE TABLE "mch_buyer_group";
TRUNCATE TABLE "pt_mail_template";
TRUNCATE TABLE "pt_mail_queue";


TRUNCATE TABLE "ws_wholesaler";
TRUNCATE TABLE "ws_rebate_rate";
TRUNCATE TABLE "ws_item";
TRUNCATE TABLE "ws_item_discount";
TRUNCATE TABLE "ws_sku_price";
TRUNCATE TABLE "ws_cart";
TRUNCATE TABLE "ws_cart_item";


TRUNCATE TABLE "pay_order";
TRUNCATE TABLE "pay_trade_data";
TRUNCATE TABLE "PAY_MERGE_ORDER";
TRUNCATE TABLE "pay_divide";


TRUNCATE TABLE "pm_coupon";
TRUNCATE TABLE "pm_coupon_bind";
TRUNCATE TABLE "pm_coupon_take";
TRUNCATE TABLE "pm_info";
TRUNCATE TABLE "pm_cash_back";


TRUNCATE TABLE "dlv_area";
TRUNCATE TABLE "dlv_coverage";
TRUNCATE TABLE "dlv_merchant_bind";


--TRUNCATE TABLE "user_role";
--TRUNCATE TABLE "user_person";
--TRUNCATE TABLE "user_credential";

TRUNCATE TABLE "pf_riseinfo";
TRUNCATE TABLE "pf_riseday";
TRUNCATE TABLE "pf_riselog";

TRUNCATE TABLE "comm_qr_template";
TRUNCATE TABLE "portal_nav";
TRUNCATE TABLE "portal_nav_type";
TRUNCATE TABLE "portal_floor_ad";
TRUNCATE TABLE "portal_floor_link";
TRUNCATE TABLE "sys_kv";

-- 2024-10-21
TRUNCATE TABLE "invoice_item";
TRUNCATE TABLE "invoice_record";
TRUNCATE TABLE "invoice_tenant";
TRUNCATE TABLE "invoice_title";

TRUNCATE TABLE "approval";
TRUNCATE TABLE "approval_flow";


TRUNCATE TABLE "chat_conversation";
TRUNCATE TABLE "chat_msg";

TRUNCATE TABLE "mch_service_order";
TRUNCATE TABLE "mch_transform_setting";
TRUNCATE TABLE "mch_staff";
TRUNCATE TABLE "mch_staff_extent";
TRUNCATE TABLE "mch_staff_transfer";
TRUNCATE TABLE "mch_bill";
TRUNCATE TABLE "mch_balance_log";


TRUNCATE TABLE "workorder";
TRUNCATE TABLE "workorder_comment";



delete from mch_merchant where id>1;

DELETE FROM mch_authenticate where mch_id NOT IN(select id from mch_merchant);
DELETE FROM mch_online_shop where vendor_id NOT IN(select id from mch_merchant);

DELETE FROM mch_account where mch_id NOT IN(select id from mch_merchant);











