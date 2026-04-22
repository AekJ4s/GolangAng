import { Component, OnInit, ChangeDetectorRef } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../../services/auth.service';
import { User } from '../../models/user.model';

@Component({
  selector: 'app-dashboard',
  imports: [],
  templateUrl: './dashboard.view.html',
  styleUrl: './dashboard.style.css',
})
export class DashboardController implements OnInit {
  user: User = { id: 0, username: '' };

  constructor(
    private auth: AuthService,
    private router: Router,
    private cdr: ChangeDetectorRef
  ) {}

  ngOnInit(): void {
    this.auth.getProfile().subscribe({
      next: (res) => {
        if (res.data) {
          this.user = res.data;
          this.cdr.detectChanges();
        }
      },
      error: () => this.router.navigate(['/login']),
    });
  }

  onLogout(): void {
    this.auth.logout();
    this.router.navigate(['/login']);
  }
}
