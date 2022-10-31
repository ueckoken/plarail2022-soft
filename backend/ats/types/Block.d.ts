export interface Block{
  id:string;
  state: BlockState;
}

type BlockState = 'ERROR' | 'OPEN' | 'CLOSE';