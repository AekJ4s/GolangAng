import { Component, ChangeDetectorRef } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from '../../services/auth.service';
import { LoginRequest } from '../../models/user.model';

@Component({
  selector: 'app-login',
  imports: [FormsModule],
  templateUrl: './login.view.html',
  styleUrl: './login.style.css',
})
export class LoginController {
  form: LoginRequest = { username: '', password: '' };
  errorMessage = '';

  constructor(
    private auth: AuthService,
    private router: Router,
    private cdr: ChangeDetectorRef
  ) {}

  onLogin(): void {
    this.errorMessage = '';

    if (!this.form.username || !this.form.password) {
      this.errorMessage = 'กรุณากรอกชื่อผู้ใช้งานและรหัสผ่าน';
      return;
    }

    this.auth.login(this.form).subscribe({
      next: (res) => {
        if (res.code === '1000') {
          this.router.navigate(['/dashboard']);
        } else {
          this.errorMessage = 'ชื่อผู้ใช้หรือรหัสผ่านไม่ถูกต้อง';
          this.cdr.detectChanges();
        }
      },
      error: () => {
        this.errorMessage = 'ชื่อผู้ใช้หรือรหัสผ่านไม่ถูกต้อง';
        this.cdr.detectChanges();
      },
    });
  }

  goToRegister(): void {
    this.router.navigate(['/register']);
  }
}
