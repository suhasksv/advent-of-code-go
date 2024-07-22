#!/bin/bash

for day in {01..25}; do
	folder_name="day-$day"
	mkdir "$folder_name"

	touch "$folder_name/input.txt"
	
	touch "$folder_name/main1.go"
    	touch "$folder_name/main2.go"
	
	touch "$folder_name/main1.lua"
	touch "$folder_name/main2.lua"

	touch "$folder_name/main1.rs"
	touch "$folder_name/main2.rs"

	touch "$folder_name/main1.py"
   	touch "$folder_name/main2.py"

	touch "$folder_name/main1.c"
    	touch "$folder_name/main2.c"

	touch "$folder_name/main1.dart"
	touch "$folder_name/main2.dart"

	touch "$folder_name/main1.ex"
	touch "$folder_name/main2.ex"
done

