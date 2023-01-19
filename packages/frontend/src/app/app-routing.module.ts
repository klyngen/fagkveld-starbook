import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginScreenComponent } from './login-screen/login-screen.component';
import { PersonListComponent } from './person-list/person-list.component';

const routes: Routes = [{
  path: '',
  component: PersonListComponent
}, {
  path: 'login',
    component: LoginScreenComponent
  }];


@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
