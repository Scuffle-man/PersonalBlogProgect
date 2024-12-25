import api from './api';

export const API_ROUTES = {
  LOGIN: '/login',
  REGISTER: '/register',
  DELETE_ACCOUNT: '/users/me'
};

export const authService = {
  async login(credentials) {
    const response = await api.post(API_ROUTES.LOGIN, credentials);
    if (response.token) {
      localStorage.setItem('token', response.token);
    }
    return response;
  },

  async register(userData) {
    return await api.post(API_ROUTES.REGISTER, userData);
  },

  logout() {
    localStorage.removeItem('token');
  }
}; 