SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema cfc
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema cfc
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `cfc` DEFAULT CHARACTER SET utf8 ;
-- -----------------------------------------------------
-- Schema cfc
-- -----------------------------------------------------
USE `cfc` ;

-- -----------------------------------------------------
-- Table `cfc`.`Person`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `cfc`.`Person` (
  `userId` INT UNSIGNED NOT NULL,
  `userName` VARCHAR(45) NOT NULL,
  `password` VARCHAR(45) NOT NULL,
  `firstName` VARCHAR(45) NOT NULL,
  `lastName` VARCHAR(45) NOT NULL,
  `address` VARCHAR(45) NOT NULL,
  `phoneNumber` VARCHAR(45) NOT NULL,
  `role` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`userId`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `cfc`.`Client`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `cfc`.`Client` (
  `clientId` INT UNSIGNED NOT NULL,
  `Person_userId` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`clientId`, `Person_userId`),
  INDEX `fk_Client_Person1_idx` (`Person_userId` ASC) VISIBLE,
  CONSTRAINT `fk_Client_Person1`
    FOREIGN KEY (`Person_userId`)
    REFERENCES `cfc`.`Person` (`userId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `cfc`.`Family_Member`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `cfc`.`Family_Member` (
  `familyId` INT UNSIGNED NOT NULL,
  `relationship` VARCHAR(45) NOT NULL,
  `Client_clientId` INT UNSIGNED NOT NULL,
  `Person_userId` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`familyId`, `Client_clientId`, `Person_userId`),
  INDEX `fk_Family_Member_Client1_idx` (`Client_clientId` ASC) VISIBLE,
  INDEX `fk_Family_Member_Person1_idx` (`Person_userId` ASC) VISIBLE,
  CONSTRAINT `fk_Family_Member_Client1`
    FOREIGN KEY (`Client_clientId`)
    REFERENCES `cfc`.`Client` (`clientId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Family_Member_Person1`
    FOREIGN KEY (`Person_userId`)
    REFERENCES `cfc`.`Person` (`userId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `cfc`.`Clinician`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `cfc`.`Clinician` (
  `clinicianId` INT UNSIGNED NOT NULL,
  `Person_userId` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`clinicianId`, `Person_userId`),
  INDEX `fk_Clinician_Person1_idx` (`Person_userId` ASC) VISIBLE,
  CONSTRAINT `fk_Clinician_Person1`
    FOREIGN KEY (`Person_userId`)
    REFERENCES `cfc`.`Person` (`userId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `cfc`.`Client_has_Clinician`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `cfc`.`Client_has_Clinician` (
  `Client_clientId` INT UNSIGNED NOT NULL,
  `Clinician_clinicianId` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`Client_clientId`, `Clinician_clinicianId`),
  INDEX `fk_Client_has_Clinician_Clinician1_idx` (`Clinician_clinicianId` ASC) VISIBLE,
  INDEX `fk_Client_has_Clinician_Client1_idx` (`Client_clientId` ASC) VISIBLE,
  CONSTRAINT `fk_Client_has_Clinician_Client1`
    FOREIGN KEY (`Client_clientId`)
    REFERENCES `cfc`.`Client` (`clientId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Client_has_Clinician_Clinician1`
    FOREIGN KEY (`Clinician_clinicianId`)
    REFERENCES `cfc`.`Clinician` (`clinicianId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `cfc`.`Support_Network`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `cfc`.`Support_Network` (
  `supportId` INT UNSIGNED NOT NULL,
  `relationship` VARCHAR(45) NULL,
  `phoneNumber` VARCHAR(45) NULL,
  `Client_clientId` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`supportId`, `Client_clientId`),
  INDEX `fk_Support_Network_Client1_idx` (`Client_clientId` ASC) VISIBLE,
  CONSTRAINT `fk_Support_Network_Client1`
    FOREIGN KEY (`Client_clientId`)
    REFERENCES `cfc`.`Client` (`clientId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `cfc`.`Safety_Plan`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `cfc`.`Safety_Plan` (
  `safetyId` INT UNSIGNED NOT NULL,
  `triggers` VARCHAR(45) NULL,
  `warningSigns` VARCHAR(45) NULL,
  `destructiveBehaviors` VARCHAR(45) NULL,
  `internalStrategies` VARCHAR(45) NULL,
  `updatedDatetime` DATETIME NULL,
  `Client_clientId` INT UNSIGNED NOT NULL,
  `Clinician_clinicianId` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`safetyId`, `Client_clientId`, `Clinician_clinicianId`),
  INDEX `fk_Safety_Plan_Client1_idx` (`Client_clientId` ASC) VISIBLE,
  INDEX `fk_Safety_Plan_Clinician1_idx` (`Clinician_clinicianId` ASC) VISIBLE,
  CONSTRAINT `fk_Safety_Plan_Client1`
    FOREIGN KEY (`Client_clientId`)
    REFERENCES `cfc`.`Client` (`clientId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Safety_Plan_Clinician1`
    FOREIGN KEY (`Clinician_clinicianId`)
    REFERENCES `cfc`.`Clinician` (`clinicianId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `cfc`.`Agencies`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `cfc`.`Agencies` (
  `agencyId` INT NOT NULL,
  `name` VARCHAR(45) NULL,
  `phoneNumber` VARCHAR(45) NULL,
  `specialization` VARCHAR(45) NULL,
  PRIMARY KEY (`agencyId`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `cfc`.`Safety_Plan_has_Agencies`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `cfc`.`Safety_Plan_has_Agencies` (
  `Safety_Plan_safetyId` INT UNSIGNED NOT NULL,
  `Agencies_agencyId` INT NOT NULL,
  PRIMARY KEY (`Safety_Plan_safetyId`, `Agencies_agencyId`),
  INDEX `fk_Safety_Plan_has_Agencies_Agencies1_idx` (`Agencies_agencyId` ASC) VISIBLE,
  INDEX `fk_Safety_Plan_has_Agencies_Safety_Plan1_idx` (`Safety_Plan_safetyId` ASC) VISIBLE,
  CONSTRAINT `fk_Safety_Plan_has_Agencies_Safety_Plan1`
    FOREIGN KEY (`Safety_Plan_safetyId`)
    REFERENCES `cfc`.`Safety_Plan` (`safetyId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Safety_Plan_has_Agencies_Agencies1`
    FOREIGN KEY (`Agencies_agencyId`)
    REFERENCES `cfc`.`Agencies` (`agencyId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `cfc`.`Appointments`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `cfc`.`Appointments` (
  `appointmentId` INT NOT NULL,
  `appointmentTime` DATETIME NULL,
  `appointmentMedium` VARCHAR(45) NULL,
  `Client_clientId` INT UNSIGNED NOT NULL,
  `Clinician_clinicianId` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`appointmentId`, `Client_clientId`, `Clinician_clinicianId`),
  INDEX `fk_Appointments_Client1_idx` (`Client_clientId` ASC) VISIBLE,
  INDEX `fk_Appointments_Clinician1_idx` (`Clinician_clinicianId` ASC) VISIBLE,
  CONSTRAINT `fk_Appointments_Client1`
    FOREIGN KEY (`Client_clientId`)
    REFERENCES `cfc`.`Client` (`clientId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Appointments_Clinician1`
    FOREIGN KEY (`Clinician_clinicianId`)
    REFERENCES `cfc`.`Clinician` (`clinicianId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
