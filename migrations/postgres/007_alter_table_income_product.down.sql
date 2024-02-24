alter table if exists income_products
     drop column if exists branch_id;


delete from branches where id = '33c81610-854b-42d0-9610-44f880f37b3b';