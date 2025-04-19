import { Component, EventEmitter, Input, Output } from '@angular/core';
import { ReferencesWithSimilarity } from '../../core/interfaces/chat.service.interface';

@Component({
  selector: 'app-refrence-traversal',
  imports: [],
  templateUrl: './refrence-traversal.component.html',
  styleUrl: './refrence-traversal.component.scss'
})
export class RefrenceTraversalComponent {
  @Input() showReferencesTraversal: boolean = false;
  @Input() references: ReferencesWithSimilarity[] = [];
  @Output() showRefernceTraversalEmitter = new EventEmitter<boolean>();
  showRefernces() {
    this.showRefernceTraversalEmitter.emit(false);
  }

}
