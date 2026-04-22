import { Routes } from '@angular/router';
import { authGuard } from './core/guards/auth.guard';

export const routes: Routes = [
  { path: '', redirectTo: 'login', pathMatch: 'full' },
  {
    path: 'login',
    loadComponent: () =>
      import('./controllers/login/login.controller').then(m => m.LoginController),
  },
  {
    path: 'register',
    loadComponent: () =>
      import('./controllers/register/register.controller').then(m => m.RegisterController),
  },
  {
    path: 'dashboard',
    loadComponent: () =>
      import('./controllers/dashboard/dashboard.controller').then(m => m.DashboardController),
    canActivate: [authGuard],
  },
];
