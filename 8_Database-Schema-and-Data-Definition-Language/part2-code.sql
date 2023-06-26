-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------
-- -----------------------------------------------------
-- Schema alta_online_shop
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema alta_online_shop
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `alta_online_shop` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci ;
USE `alta_online_shop` ;

-- -----------------------------------------------------
-- Table `alta_online_shop`.`users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `alta_online_shop`.`users` (
  `user_id` INT NOT NULL,
  `name` VARCHAR(45) NOT NULL,
  `address` VARCHAR(255) NULL DEFAULT NULL,
  `date_of_birth` DATE NULL DEFAULT NULL,
  `status_user` VARCHAR(45) NULL DEFAULT NULL,
  `gender` VARCHAR(45) NULL DEFAULT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`user_id`),
  UNIQUE INDEX `user_id_UNIQUE` (`user_id` ASC) VISIBLE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `alta_online_shop`.`address`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `alta_online_shop`.`address` (
  `address_id` INT NOT NULL,
  `user_id` INT NOT NULL,
  `alamat` LONGTEXT NOT NULL,
  PRIMARY KEY (`address_id`),
  UNIQUE INDEX `address_id_UNIQUE` (`address_id` ASC) VISIBLE,
  UNIQUE INDEX `user_id_UNIQUE` (`user_id` ASC) VISIBLE,
  CONSTRAINT `iduser-address`
    FOREIGN KEY (`user_id`)
    REFERENCES `alta_online_shop`.`users` (`user_id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `alta_online_shop`.`payment_method`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `alta_online_shop`.`payment_method` (
  `pm_id` INT NOT NULL,
  `method_name` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`pm_id`),
  UNIQUE INDEX `pm_id_UNIQUE` (`pm_id` ASC) VISIBLE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `alta_online_shop`.`payment_method_description`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `alta_online_shop`.`payment_method_description` (
  `pmd_id` INT NOT NULL,
  `description` LONGTEXT NOT NULL,
  `pm_id` INT NOT NULL,
  PRIMARY KEY (`pmd_id`, `pm_id`),
  UNIQUE INDEX `pmd_id_UNIQUE` (`pmd_id` ASC) VISIBLE,
  INDEX `fk_payment_method_description_payment_method1_idx` (`pm_id` ASC) VISIBLE,
  CONSTRAINT `fk_payment_method_description_payment_method1`
    FOREIGN KEY (`pm_id`)
    REFERENCES `alta_online_shop`.`payment_method` (`pm_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `alta_online_shop`.`products`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `alta_online_shop`.`products` (
  `product_id` INT NOT NULL,
  `product_name` VARCHAR(255) NOT NULL,
  `product_type` VARCHAR(45) NOT NULL,
  `product_description` LONGTEXT NULL DEFAULT NULL,
  `oprator_name` VARCHAR(255) NOT NULL,
  `pm_id` INT NOT NULL,
  PRIMARY KEY (`product_id`),
  UNIQUE INDEX `product_id_UNIQUE` (`product_id` ASC) VISIBLE,
  UNIQUE INDEX `pm_id_UNIQUE` (`pm_id` ASC) VISIBLE,
  CONSTRAINT `paymentm_id`
    FOREIGN KEY (`pm_id`)
    REFERENCES `alta_online_shop`.`payment_method` (`pm_id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `alta_online_shop`.`transactions`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `alta_online_shop`.`transactions` (
  `transaction_id` INT NOT NULL,
  `user_id` INT NOT NULL,
  `product_id` INT NOT NULL,
  `transaction_date` DATETIME NOT NULL,
  `total_amount` DECIMAL(10,0) NOT NULL,
  PRIMARY KEY (`transaction_id`),
  UNIQUE INDEX `transaction_id_UNIQUE` (`transaction_id` ASC) VISIBLE,
  UNIQUE INDEX `user_id_UNIQUE` (`user_id` ASC) VISIBLE,
  UNIQUE INDEX `product_id_UNIQUE` (`product_id` ASC) VISIBLE,
  CONSTRAINT `product_id`
    FOREIGN KEY (`product_id`)
    REFERENCES `alta_online_shop`.`products` (`product_id`),
  CONSTRAINT `user_id`
    FOREIGN KEY (`user_id`)
    REFERENCES `alta_online_shop`.`users` (`user_id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `alta_online_shop`.`user_payment_method_detail`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `alta_online_shop`.`user_payment_method_detail` (
  `upmd_id` INT NOT NULL,
  `user_id` INT NOT NULL,
  `pm_id` INT NOT NULL,
  PRIMARY KEY (`upmd_id`),
  UNIQUE INDEX `upmd_id_UNIQUE` (`upmd_id` ASC) VISIBLE,
  UNIQUE INDEX `user_id_UNIQUE` (`user_id` ASC) VISIBLE,
  UNIQUE INDEX `pm_id_UNIQUE` (`pm_id` ASC) VISIBLE,
  CONSTRAINT `pmid-upmd`
    FOREIGN KEY (`pm_id`)
    REFERENCES `alta_online_shop`.`payment_method` (`pm_id`),
  CONSTRAINT `userid-upmd`
    FOREIGN KEY (`user_id`)
    REFERENCES `alta_online_shop`.`users` (`user_id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
