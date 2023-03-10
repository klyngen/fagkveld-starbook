import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';
import { map } from 'rxjs';
import { PersonService } from '../person.service';

@Component({
  selector: 'app-create-person',
  templateUrl: './create-person.component.html',
  styleUrls: ['./create-person.component.scss']
})
export class CreatePersonComponent {

  constructor(private personService: PersonService) { }

  nameControl = new UntypedFormControl("");
  pictureControl = new UntypedFormControl("");

  isValidForm = this.nameControl.valueChanges.pipe(map(value => value?.length > 3));

  onSubmitClick() {
    this.personService.createPerson(this.nameControl.value, this.pictureControl.value).subscribe();
    this.nameControl.setValue("");
    this.pictureControl.setValue("");
  }

}
