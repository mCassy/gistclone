package dto

type FileDTO struct{
	Name    string `json:"name" example:"class"`
	Content string `json:"content" example:"<?php echo 1;?>"`
}