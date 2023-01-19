import { NgModule } from '@angular/core';
import { AuthModule } from 'angular-auth-oidc-client';

@NgModule({
  imports: [
    AuthModule.forRoot({
      config: {
        authority:
          'https://login.microsoftonline.com/76749190-4427-4b08-a3e4-161767dd1b73/v2.0',
        secureRoutes: ['http://localhost:1337', 'http://127.0.0.1:1337/'],
        authWellknownEndpointUrl:
          'https://login.microsoftonline.com/76749190-4427-4b08-a3e4-161767dd1b73/v2.0',
        redirectUrl: window.location.origin,
        clientId: '5ad0b180-b2ad-43e8-bc4c-fba59a4b5108',
        scope:
          'openid profile api://5ad0b180-b2ad-43e8-bc4c-fba59a4b5108/access_as_user',
        responseType: 'code',
        silentRenew: true,
        useRefreshToken: true,
        maxIdTokenIatOffsetAllowedInSeconds: 600,
        issValidationOff: false,
        autoUserInfo: false,
        customParamsAuthRequest: {
          prompt: 'select_account', // login, consent
        },
      },
    }),
  ],
  exports: [AuthModule],
})
export class AuthConfigModule {}
