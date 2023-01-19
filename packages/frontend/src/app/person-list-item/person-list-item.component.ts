import { Component, Input } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';
import { map } from 'rxjs';
import { Person } from '../models/person';
import { StarService } from '../star.service';

@Component({
  selector: 'app-person-list-item',
  templateUrl: './person-list-item.component.html',
  styleUrls: ['./person-list-item.component.scss']
})
export class PersonListItemComponent {

  @Input()
  person: Person | undefined;

  textField = new UntypedFormControl("");

  canSubmit$ = this.textField.valueChanges.pipe(map(value => value?.length > 3));


  constructor(private starService: StarService) {
  }


  onSubmitClick() {
    if (this.person) {
      this.starService.addStar(this.textField.value, this.person.ID)
      this.textField.setValue("");
    }
  }
}
