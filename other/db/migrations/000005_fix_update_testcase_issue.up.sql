ALTER TABLE `submission_result`
  DROP CONSTRAINT `submission_result_ibfk_2`;
ALTER TABLE `submission_result`
  ADD CONSTRAINT `submission_result_ibfk_2` FOREIGN KEY (`testcase_id`) REFERENCES `testcase`(`id`);

ALTER TABLE `testcase` ADD `revision` INT NOT NULL AFTER `assignment_id`;
