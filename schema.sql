-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema expense
-- -----------------------------------------------------
DROP SCHEMA IF EXISTS `expense` ;

-- -----------------------------------------------------
-- Schema expense
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `expense` DEFAULT CHARACTER SET utf8 ;
SHOW WARNINGS;
USE `expense` ;

-- -----------------------------------------------------
-- Table `expense`.`user`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `expense`.`user` ;

SHOW WARNINGS;
CREATE TABLE IF NOT EXISTS `expense`.`user` (
  `user_id` INT NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(16) NOT NULL,
  `email` VARCHAR(255) NULL,
  `password` VARCHAR(256) NOT NULL,
  `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`));

SHOW WARNINGS;

-- -----------------------------------------------------
-- Table `expense`.`budget`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `expense`.`budget` ;

SHOW WARNINGS;
CREATE TABLE IF NOT EXISTS `expense`.`budget` (
  `budget_id` INT NOT NULL AUTO_INCREMENT,
  `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` TIMESTAMP NULL,
  `name` MEDIUMTEXT NOT NULL,
  `description` LONGTEXT NULL,
  `amount` DECIMAL(13,4) NOT NULL DEFAULT 0.0000,
  `end_date` DATE NOT NULL,
  `start_date` DATE NOT NULL,
  `deleted` TINYINT NULL DEFAULT 0,
  `created_by` INT NOT NULL,
  PRIMARY KEY (`budget_id`, `created_by`),
  CONSTRAINT `fk_budget_user`
    FOREIGN KEY (`created_by`)
    REFERENCES `expense`.`user` (`user_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

SHOW WARNINGS;
CREATE INDEX `fk_budget_user_idx` ON `expense`.`budget` (`created_by` ASC);

SHOW WARNINGS;

-- -----------------------------------------------------
-- Table `expense`.`category`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `expense`.`category` ;

SHOW WARNINGS;
CREATE TABLE IF NOT EXISTS `expense`.`category` (
  `category_id` INT NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`category_id`));

SHOW WARNINGS;

-- -----------------------------------------------------
-- Table `expense`.`expense`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `expense`.`expense` ;

SHOW WARNINGS;
CREATE TABLE IF NOT EXISTS `expense`.`expense` (
  `expense_id` INT NOT NULL AUTO_INCREMENT,
  `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` TIMESTAMP NULL,
  `name` MEDIUMTEXT NOT NULL,
  `description` LONGTEXT NOT NULL,
  `amount` DECIMAL(13,4) NOT NULL DEFAULT 0.0000,
  `date` DATE NOT NULL,
  `deleted` TINYINT NULL DEFAULT 0,
  `budget` INT NOT NULL,
  `category` INT NOT NULL,
  PRIMARY KEY (`expense_id`, `budget`, `category`),
  CONSTRAINT `fk_budget_id`
    FOREIGN KEY (`budget`)
    REFERENCES `expense`.`budget` (`budget_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_category_id`
    FOREIGN KEY (`category`)
    REFERENCES `expense`.`category` (`category_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);

SHOW WARNINGS;
CREATE INDEX `fk_budget_id_idx` ON `expense`.`expense` (`budget` ASC);

SHOW WARNINGS;
CREATE INDEX `fk_category_id_idx` ON `expense`.`expense` (`category` ASC);

SHOW WARNINGS;
USE `expense` ;

-- -----------------------------------------------------
-- Placeholder table for view `expense`.`vw_budget`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `expense`.`vw_budget` (`budget_id` INT, `name` INT, `description` INT, `amount` INT, `start_date` INT, `end_date` INT);
SHOW WARNINGS;

-- -----------------------------------------------------
-- Placeholder table for view `expense`.`vw_expense`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `expense`.`vw_expense` (`expense_id` INT, `name` INT, `description` INT, `amount` INT, `date` INT, `budget` INT, `category` INT);
SHOW WARNINGS;

-- -----------------------------------------------------
-- procedure budget_create
-- -----------------------------------------------------

USE `expense`;
DROP procedure IF EXISTS `expense`.`budget_create`;
SHOW WARNINGS;

DELIMITER $$
USE `expense`$$
CREATE PROCEDURE budget_create
(IN iname MEDIUMTEXT, IN idescription LONGTEXT, IN iamount DECIMAL(13,4), IN istart_date TEXT, IN iend_date TEXT, OUT oid INT)
BEGIN
	INSERT INTO `budget` (
    update_time,
    name,
    description,
    amount,
    end_date,
    start_date
    ) VALUES (
		NOW(),
		iname,
        idescription,
        iamount,
        istart_date,
        iend_date
    );
    SET oid = LAST_INSERT_ID();
END$$

DELIMITER ;
SHOW WARNINGS;

-- -----------------------------------------------------
-- procedure sp_budget_create
-- -----------------------------------------------------

USE `expense`;
DROP procedure IF EXISTS `expense`.`sp_budget_create`;
SHOW WARNINGS;

DELIMITER $$
USE `expense`$$
CREATE PROCEDURE sp_budget_create
(IN iname MEDIUMTEXT, IN idescription LONGTEXT, IN iamount DECIMAL(13,4), IN istart_date TEXT, IN iend_date TEXT, OUT oid INT)
BEGIN
	INSERT INTO `budget` (
    update_time,
    name,
    description,
    amount,
    end_date,
    start_date
    ) VALUES (
		NOW(),
		iname,
        idescription,
        iamount,
        istart_date,
        iend_date
    );
    SET oid = LAST_INSERT_ID();
END$$

DELIMITER ;
SHOW WARNINGS;

-- -----------------------------------------------------
-- procedure sp_expense_create
-- -----------------------------------------------------

USE `expense`;
DROP procedure IF EXISTS `expense`.`sp_expense_create`;
SHOW WARNINGS;

DELIMITER $$
USE `expense`$$
CREATE PROCEDURE sp_expense_create
(IN iname MEDIUMTEXT, IN idescription LONGTEXT, IN iamount DECIMAL(13,4), IN idate TEXT, IN ibudget_id INT, IN icategory_id INT, OUT oid INT)
BEGIN
	INSERT INTO `budget` (
    update_time,
    name,
    description,
    amount,
    date,
    budget,
    category
    ) VALUES (
		NOW(),
		iname,
        idescription,
        iamount,
        idate,
        ibudget_id,
        icategory_id
    );
    SET oid = LAST_INSERT_ID();
END$$

DELIMITER ;
SHOW WARNINGS;

-- -----------------------------------------------------
-- procedure sp_budget_update
-- -----------------------------------------------------

USE `expense`;
DROP procedure IF EXISTS `expense`.`sp_budget_update`;
SHOW WARNINGS;

DELIMITER $$
USE `expense`$$
CREATE PROCEDURE sp_budget_update
(IN id INT, IN iname MEDIUMTEXT, IN idescription LONGTEXT, IN iamount DECIMAL(13,4), IN istart_date TEXT, IN iend_date TEXT, IN iuser_id INT, OUT oid INT)
BEGIN
	IF id = 0 THEN
		INSERT INTO `budget` (
		update_time,
		name,
		description,
		amount,
		end_date,
		start_date,
        created_by
		) VALUES (
			NOW(),
			iname,
			idescription,
			iamount,
			istart_date,
			iend_date,
            iuser_id
		);
		SET oid = LAST_INSERT_ID();
	ELSE
		UPDATE `budget` SET
		update_time = NOW(),
        name = iname,
        description = idescription,
        amount = iamount,
        start_date = istart_date,
        end_date = iend_date,
        created_by = iuser_id 
        WHERE budget_id = id;
    END IF;
END$$

DELIMITER ;
SHOW WARNINGS;

-- -----------------------------------------------------
-- procedure sp_expense_update
-- -----------------------------------------------------

USE `expense`;
DROP procedure IF EXISTS `expense`.`sp_expense_update`;
SHOW WARNINGS;

DELIMITER $$
USE `expense`$$
CREATE PROCEDURE sp_expense_update
(IN id INT, IN iname MEDIUMTEXT, IN idescription LONGTEXT, IN iamount DECIMAL(13,4), IN idate TEXT, IN ibudget_id INT, IN icategory_id INT, OUT oid INT)
BEGIN
	IF id = 0 THEN
		INSERT INTO `budget` (
		update_time,
		name,
		description,
		amount,
		date,
		budget,
		category
		) VALUES (
			NOW(),
			iname,
			idescription,
			iamount,
			idate,
			ibudget_id,
			icategory_id
		);
		SET oid = LAST_INSERT_ID();
	ELSE
		UPDATE `expense` SET
        update_time = NOW(),
        name = iname,
        description = idescription,
        amount = iamount,
        date = idate,
        budget = ibudget_id,
        category = icategory_id
        WHERE expense_id = id;
    END IF;
END$$

DELIMITER ;
SHOW WARNINGS;

-- -----------------------------------------------------
-- procedure sp_delete_budget
-- -----------------------------------------------------

USE `expense`;
DROP procedure IF EXISTS `expense`.`sp_delete_budget`;
SHOW WARNINGS;

DELIMITER $$
USE `expense`$$
CREATE PROCEDURE sp_delete_budget(IN id INT)
BEGIN
	UPDATE `budget`
    SET deleted = 1
    WHERE budget_id = id;
END$$

DELIMITER ;
SHOW WARNINGS;

-- -----------------------------------------------------
-- procedure sp_delete_expense
-- -----------------------------------------------------

USE `expense`;
DROP procedure IF EXISTS `expense`.`sp_delete_expense`;
SHOW WARNINGS;

DELIMITER $$
USE `expense`$$
CREATE PROCEDURE sp_delete_expense(IN id INT)
BEGIN
	UPDATE `expense`
    SET deleted = 1
    WHERE expense_id = id;
END$$

DELIMITER ;
SHOW WARNINGS;

-- -----------------------------------------------------
-- View `expense`.`vw_budget`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `expense`.`vw_budget`;
SHOW WARNINGS;
DROP VIEW IF EXISTS `expense`.`vw_budget` ;
SHOW WARNINGS;
USE `expense`;
CREATE  OR REPLACE VIEW `vw_budget` AS
SELECT budget_id, name, description, amount, start_date, end_date FROM
`budget` WHERE deleted <> 1;
SHOW WARNINGS;

-- -----------------------------------------------------
-- View `expense`.`vw_expense`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `expense`.`vw_expense`;
SHOW WARNINGS;
DROP VIEW IF EXISTS `expense`.`vw_expense` ;
SHOW WARNINGS;
USE `expense`;
CREATE  OR REPLACE VIEW `vw_expense` AS
SELECT expense_id, name, description, amount, date, budget, category FROM
`expense`
WHERE deleted <> 1;
SHOW WARNINGS;

SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
