<?php

/**
 * MetaChat server core
 * 
 * TASK:
 * user->register/login/logout
 * group->add_member/remove_member/create_group/remove_group
 * admin->dashboard/database/users/groups
 * 
 * And test modified codes && rewrite android client
 */

class DB
{
    private $db_path = './db.json';
    public $data = [];

    public function __construct()
    {
        $this->data = json_decode(file_get_contents($this->db_path), true);
    }

    public function getData($type = 'array')
    {
        switch ($type)
        {
            case 'array':
                return $this->data;

            case 'json':
                return json_encode($this->data);
        }
    }

    public function saveData()
    {
        file_put_contents($this->db_path, json_encode($this->data), LOCK_EX);
    }
}

/**
 * main
 */

$db = new DB;

/* bug fix: only process first param because other params will be processed in the same routine */
foreach (($api = $_GET) as $key => $value)
{
    switch ($key)
    {
        case 'admin':
            break;
        default:
            echo json_encode(array('err' => 'no_such_api'));
            break;
    }
    break;
}
