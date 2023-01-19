import { Component } from '@angular/core';
import { OidcSecurityService } from 'angular-auth-oidc-client';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent {
  title = 'tarbook-frontenddddddsssss';

  constructor(oidcSecurityService: OidcSecurityService) {
    oidcSecurityService.checkAuth().subscribe();
  }
}
