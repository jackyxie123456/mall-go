


## 修改表

### 修改表 

- pms_brand
  新增 两个字段 
  - name_en  品牌英文名
  - brand_story_en  品牌故事英文 
- pms_product_category 
  新增四个字段 , 商品筛选属性，是做什么用的？
  - name_en 
  - unit_en 单位英文
  - keyword_en 
  - description_en  
- pms_product 表
  新增五个字段
  - name_en
  - sub_title_en
  - brand_name_en
  - description_en
  - product_category_name_en
- pms_product_attribute_value 表
  新增一个字段
  - value_en
- pms_product_attribute 表
  新增两个字段
  - name_en
  - input_list_en

- pms_product_attribute_category 表
  新增一个字段
  - name_en


- ums_member_level 表
  新增一个字段
  - name_en


- ums_member_receive_address
  新增一个字段
  - country

- pms_member_price
  新增一个字段
  - member_level_name_en

- sms_coupon_product_relation
  新增一个字段
  - product_name_en

- sms_coupon_product_category_relation
  新增两个字段
  - product_category_name_en
  - parent_category_name_en 

- oms_order_item
  新增两个字段
  - product_name_en
  - product_brand_en 
  - promotion_name_en
  - product_attr_en  


- oms_order_return_reason
  增加一个字段
  - name_en 

- oms_order_return_apply
  增加 五 个字段
  - product_name_en
  - product_brand_en 
  - product_attr_en  
  - reason_en
  - description_en 


### 修改表字段 

~~~shell
ALTER TABLE pms_product_attribute
ADD COLUMN `name_en` VARCHAR(255) DEFAULT '' COMMENT 'en name ',
ADD COLUMN `input_list_en` VARCHAR(255) DEFAULT '' COMMENT 'en input_list';



ALTER TABLE pms_product_attribute_value
ADD COLUMN `value_en` VARCHAR(512) DEFAULT '' COMMENT 'en value';

ALTER TABLE pms_brand
ADD COLUMN `name_en` VARCHAR(255) DEFAULT '' COMMENT 'en value',
ADD COLUMN `brand_story_en` TEXT DEFAULT '' COMMENT 'en value';



ALTER TABLE pms_product
ADD COLUMN `name_en` VARCHAR(255) DEFAULT '' COMMENT 'en name',
ADD COLUMN `sub_title_en` VARCHAR(512) DEFAULT '' COMMENT 'en sub_title',
ADD COLUMN `brand_name_en` VARCHAR(255) DEFAULT '' COMMENT 'en brand_name',
ADD COLUMN `description_en` VARCHAR(255) DEFAULT '' COMMENT 'en description_en',
ADD COLUMN `product_category_name_en` VARCHAR(255) DEFAULT '' COMMENT 'en product_category_name_en';

ALTER  TABLE pms_product_category
ADD COLUMN `name_en` VARCHAR(255) DEFAULT '' COMMENT 'en name',
ADD COLUMN `product_unit_en` VARCHAR(512) DEFAULT '' COMMENT 'en product_unit',
ADD COLUMN `keywords_en` VARCHAR(255) DEFAULT '' COMMENT 'en keywords',
ADD COLUMN `description_en` VARCHAR(255) DEFAULT '' COMMENT 'en description';


ALTER  TABLE oms_cart_item
ADD COLUMN `product_name_en` VARCHAR(255) DEFAULT '' COMMENT 'en product_name_en',
ADD COLUMN `product_sub_title_en` VARCHAR(512) DEFAULT '' COMMENT 'en product_sub_title',
ADD COLUMN `product_brand_en` VARCHAR(512) DEFAULT '' COMMENT 'en product_brand',
ADD COLUMN `product_attr_en` VARCHAR(512) DEFAULT '' COMMENT 'en product_attr';


ALTER  TABLE oms_order
ADD COLUMN `promotion_info_en` VARCHAR(255) DEFAULT '' COMMENT 'en promotion_info';

ALTER  TABLE oms_order_item
ADD COLUMN `product_name_en` VARCHAR(255) DEFAULT '' COMMENT 'en product_name',
ADD COLUMN `product_brand_en` VARCHAR(255) DEFAULT '' COMMENT 'en product_brand',
ADD COLUMN `promotion_name_en` VARCHAR(255) DEFAULT '' COMMENT 'en promotion_name_en',
ADD COLUMN `product_attr_en` VARCHAR(255) DEFAULT '' COMMENT 'en product_attr';


ALTER  TABLE pms_product_attribute_category
ADD COLUMN `name_en` VARCHAR(255) DEFAULT '' COMMENT 'en name';


ALTER  TABLE ums_member_level
ADD COLUMN `name_en` VARCHAR(64) DEFAULT '' COMMENT 'en name';

ALTER  TABLE ums_member_receive_address
ADD COLUMN `country` VARCHAR(128) DEFAULT '' COMMENT 'country';

ALTER  TABLE pms_member_price
ADD COLUMN `member_level_name_en` VARCHAR(64) DEFAULT '' COMMENT 'member_level_name en';

ALTER  TABLE sms_coupon_product_relation
ADD COLUMN `product_name_en` VARCHAR(500) DEFAULT '' COMMENT 'product_name en';


ALTER  TABLE sms_coupon_product_category_relation
ADD COLUMN `product_category_name_en` VARCHAR(200) DEFAULT '' COMMENT 'product_category_name_ en',
ADD COLUMN `parent_category_name_en` VARCHAR(200) DEFAULT '' COMMENT 'parent_category_name en';



ALTER  TABLE oms_order_return_reason
ADD COLUMN `name_en` VARCHAR(200) DEFAULT '' COMMENT 'name_en en';


ALTER  TABLE oms_order_return_apply
ADD COLUMN `product_name_en` VARCHAR(200) DEFAULT '' COMMENT 'product_name_en en',
ADD COLUMN `product_brand_en` VARCHAR(200) DEFAULT '' COMMENT 'product_brand_en en',
ADD COLUMN `product_attr_en` VARCHAR(200) DEFAULT '' COMMENT 'product_attr_en en',
ADD COLUMN `reason_en` VARCHAR(200) DEFAULT '' COMMENT 'reason_en en',
ADD COLUMN `description_en` VARCHAR(200) DEFAULT '' COMMENT 'description_en en';

### 修改记录 

~~~shell

UPDATE pms_brand
JOIN pms_brand_en ON pms_brand.id = pms_brand_en.id
SET pms_brand.name_en = pms_brand_en.name,
    pms_brand.brand_story_en = pms_brand_en.brand_story
WHERE 1=1;


UPDATE pms_product_category
JOIN pms_product_category_en ON pms_product_category.id = pms_product_category_en.id
SET pms_product_category.name_en = pms_product_category_en.name,
    pms_product_category.product_unit_en = pms_product_category_en.product_unit,
    pms_product_category.keywords_en = pms_product_category_en.keywords,
    pms_product_category.description_en = pms_product_category_en.description
WHERE 1=1;

UPDATE pms_product_attribute_value
JOIN pms_product_attribute_value_en ON pms_product_attribute_value.id = pms_product_attribute_value_en.id
SET pms_product_attribute_value.value_en = pms_product_attribute_value_en.value
WHERE 1=1;



UPDATE pms_product_attribute
JOIN pms_product_attribute_en ON pms_product_attribute.id = pms_product_attribute_en.id
SET pms_product_attribute.name_en = pms_product_attribute_en.name,
    pms_product_attribute.input_list_en = pms_product_attribute_en.input_list;




UPDATE pms_product
JOIN pms_product_en ON pms_product.id = pms_product_en.id
SET pms_product.name_en = pms_product_en.name,
    pms_product.sub_title_en = pms_product_en.sub_title,
    pms_product.brand_name_en = pms_product_en.brand_name,
    pms_product.description_en = pms_product_en.description,
    pms_product.product_category_name_en = pms_product_en.product_category_name;


UPDATE pms_product_attribute_category
JOIN pms_product_en ON pms_product.id = pms_product_en.id
SET pms_product.name_en = pms_product_en.name,
    pms_product.sub_title_en = pms_product_en.sub_title,
    pms_product.brand_name_en = pms_product_en.brand_name,
    pms_product.description_en = pms_product_en.description,
    pms_product.product_category_name_en = pms_product_en.product_category_name;


UPDATE pms_member_price
JOIN ums_member_level ON pms_member_price.member_level_id = ums_member_level.id
SET pms_member_price.member_level_name_en = ums_member_level.name_en;


UPDATE oms_order_item
JOIN pms_product ON oms_order_item.product_id = pms_product.id
SET oms_order_item.product_name_en = pms_product.name_en,
    oms_order_item.product_brand_en = pms_product.brand_name_en;


UPDATE oms_cart_item
JOIN pms_product ON oms_cart_item.product_id = pms_product.id
SET oms_cart_item.product_name_en = pms_product.name_en,
    oms_cart_item.product_sub_title_en = pms_product.sub_title_en,
    oms_cart_item.product_brand_en = pms_product.brand_name_en;



~~~


### local 对应表输出

~~~shell

brand 表

Select("b.id,b.first_letter,b.sort,  b.factory_status,  b.show_status,  b.product_count,  b.product_comment_count,  b.logo, b.big_pic, b.brand_story,  b.name_en as name, b.brand_story_en as brand_story").

Select("b.*, b.name_en as name, b.brand_story_en as brand_story").


product 表

Select("p.*, p.name_en as name, p.sub_title_en as sub_title, p.brand_name_en as brand_name,p.description_en as description, p.product_category_name_en as product_category_name).



~~~