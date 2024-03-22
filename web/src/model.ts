export interface Notification {
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string;
    ID: number;
    title: string;
    author: string;
    content: string;
}

export interface Contest {
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string;
    ID: number;
    title: string;
    problems: Problem[];
    description: string;
    start_time: string;
    end_time: string;
}

export interface Problem {
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string;
    ID: number;
    title: string;
    difficulty: number;
    description: string;
    inputs: string[];
    outputs: string[];
    ranks: number[];
    time_limit: number;
}

export interface Submission {
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string;
    ID: number;
    problem_id: number;
    problem: Problem;
    user_id: number;
    language: string;
    code: string;
    status: string;
    information: string[];
    time: number;
    memory: number;
}

export interface User {
    CreatedAt?: string;
    UpdatedAt?: string;
    DeletedAt?: string;
    ID?: number;
    username?: string;
    password?: string;
    email?: string;
    personal_info?: PersonalInfo;
    account_info?: AccountInfo;
    submissions?: Submission[];
    ranks?: Rank[];
}

export interface PersonalInfo {
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string;
    user_id: number;
    first_name: string;
    last_name: string;
    gender: string;
    birthday: string;
    phone_number: string;
    introduction: string;
}

export interface AccountInfo {
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string;
    user_id: number;
    permission: number;
    is_active: boolean;
    tags: string[];
}

export interface Rank {
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string;
    user_id: number;
    contest_id: number;
    rank: number;
    group_id: number;
    class_rank: number;
    total_users: number;
}

export interface Pagination {
    pageNum: number;
    pageSize: number;
    total: number;
}