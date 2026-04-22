export interface User {
  id: number;
  username: string;
}

export interface LoginRequest {
  username: string;
  password: string;
}

export interface RegisterRequest {
  username: string;
  password: string;
}

export interface ApiResponse<T> {
  code: string;
  message: string;
  time: string;
  data?: T;
}

export interface LoginData {
  token: string;
}
