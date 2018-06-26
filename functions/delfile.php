<?php
	include_once("../config.php");
	include_once("./zdir.class.php");

	

	@$password = $_POST['password'];
	@$filepath = $_POST['filepath'];

	$zdir->delfile($password,$config,$filepath);
?>