<?php
	include_once("../config.php");
	include_once("./zdir.class.php");

	$content = $zdir->viewpdf($filepath);
	Header("Content-type: application/pdf");
	echo $content;
?>