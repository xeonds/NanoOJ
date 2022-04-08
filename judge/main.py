#!/usr/bin/python3

import pymysql as sql
from urllib import request as req
from urllib import parse

class dbQueueProc:
    # initialize connection with server
    def __init__(self, addr = ('www.jiujiuer.xyz', 'root', '1672', 'oj')):
        self.__addr = addr
        self.__conn = sql.connect(host=self.__addr[0], user=self.__addr[1], password=self.__addr[2], database=self.__addr[3])
    # check whether connection is closed and reconnect
    def __reconn(self):
        try:
            dbQueueProc.__conn.ping(False)
        except:
            self.__conn = sql.connect(host=self.__addr[0], user=self.__addr[1], password=self.__addr[2], database=self.__addr[3])
    # commit judge result
    def __set_result(self, solution_id, result):
        self.__reconn()
        try:
            self.__conn.cursor().execute('UPDATE solution SET result="'+result+'",time=0,memory=0,judgetime=NOW() WHERE solution_id="'+solution_id+'" and result<2 LIMIT 1;')
            self.__conn.commit()
        except:
            print('UPDATE solution SET result="'+result+'",time=0,memory=0,judgetime=NOW() WHERE solution_id="'+solution_id+'" and result<2 LIMIT 1;')
            # 发生错误时回滚
            self.__conn.rollback()
    # get code queue which is being processed
    def get_solution_queue(self):
        self.__reconn()
        cursor=self.__conn.cursor()
        if cursor.execute('select * from solution;'):
            self.__queue=cursor.fetchall()
    # process queue
    def proc_solution_queue(self, fun):
        self.__reconn()
        cursor=self.__conn.cursor()
        for row in self.__queue:
            if row[7]==1:
                cursor.execute('SELECT source from source_code where solution_id='+str(row[0])+';')
                self.__set_result(str(row[0]),str(fun(cursor.fetchone()[0], row[8])))

def code_judge(code, lang):
    # return value: 0/1=wait, 4=accept, ...
    lang=['C','C++','Java','Python','PHP','C#','JavaScript','Go','SQL'][lang]
    if lang=='Python':
        for row in dataset:
            if not row[0]==eval(code,row[1]):
                return 2
    return 4

queue = dbQueueProc()
queue.get_solution_queue()
queue.proc_solution_queue(code_judge)
