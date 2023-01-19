import { Component, OnInit } from '@angular/core';
import { OidcSecurityService } from 'angular-auth-oidc-client';

@Component({
  selector: 'app-login-screen',
  templateUrl: './login-screen.component.html',
  styleUrls: ['./login-screen.component.scss']
})
export class LoginScreenComponent implements OnInit {

  constructor(private readonly authenticationService: OidcSecurityService) { }

  ngOnInit(): void {
  }

  onLoginClick() {
    this.authenticationService.authorize();
  }
}
