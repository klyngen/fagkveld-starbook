import { NgModule } from '@angular/core';
import { AuthModule } from 'angular-auth-oidc-client';


@NgModule({
    imports: [AuthModule.forRoot({
        config: {
            authority: 'https://login.microsoftonline.com/76749190-4427-4b08-a3e4-161767dd1b73/v2.0',
            authWellknownEndpointUrl: 'https://login.microsoftonline.com/76749190-4427-4b08-a3e4-161767dd1b73/v2.0',
            redirectUrl: window.location.origin,
            clientId: '769bc88b-2daa-43f4-b8b1-da9505a4e80b',
            scope: 'openid profile', // 'openid profile offline_access ' + your scopes
            responseType: 'code',
            silentRenew: true,
            useRefreshToken: true,
            maxIdTokenIatOffsetAllowedInSeconds: 600,
            issValidationOff: false,
            autoUserInfo: false,
            customParamsAuthRequest: {
              prompt: 'select_account', // login, consent
            },
    }
      })],
    exports: [AuthModule],
})
export class AuthConfigModule {}
